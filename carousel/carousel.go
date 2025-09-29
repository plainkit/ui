package carousel

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/plainkit/html"
	"github.com/plainkit/icons/lucide"
	"github.com/plainkit/ui/internal/classnames"
)

type Props struct {
	ID       string
	Class    string
	Attrs    []html.Global
	Autoplay bool
	Interval int
	Loop     bool
}

type ContentProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type ItemProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type PreviousProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type NextProps struct {
	ID    string
	Class string
	Attrs []html.Global
}

type IndicatorsProps struct {
	ID    string
	Class string
	Attrs []html.Global
	Count int
}

// Carousel renders a carousel container for sliding content
func Carousel(props Props, args ...html.DivArg) html.Node {
	id := props.ID
	if id == "" {
		id = randomID("carousel")
	}

	interval := props.Interval
	if interval == 0 {
		interval = 5000
	}

	divArgs := []html.DivArg{
		html.AId(id),
		html.AClass(classnames.Merge("relative overflow-hidden w-full", props.Class)),
		html.AData("pui-carousel", ""),
		html.AData("pui-carousel-current", "0"),
		html.AData("pui-carousel-autoplay", strconv.FormatBool(props.Autoplay)),
		html.AData("pui-carousel-interval", fmt.Sprintf("%d", interval)),
		html.AData("pui-carousel-loop", strconv.FormatBool(props.Loop)),
	}

	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...).WithAssets("", carouselJS, "ui-carousel")
}

// Content creates the carousel track that contains the items
func Content(props ContentProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("flex h-full w-full transition-transform duration-500 ease-in-out cursor-grab", props.Class)),
		html.AData("pui-carousel-track", ""),
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

// Item creates a carousel item slide
func Item(props ItemProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("flex-shrink-0 w-full h-full relative", props.Class)),
		html.AData("pui-carousel-item", ""),
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

// Previous creates a previous navigation button
func Previous(props PreviousProps, args ...html.ButtonArg) html.Node {
	buttonArgs := []html.ButtonArg{
		html.AClass(classnames.Merge("absolute left-2 top-1/2 transform -translate-y-1/2 p-2 rounded-full bg-black/20 text-white hover:bg-black/40 focus:outline-none", props.Class)),
		html.AData("pui-carousel-prev", ""),
		html.AAria("label", "Previous slide"),
		html.AType("button"),
		lucide.ChevronLeft(html.AClass("h-4 w-4")),
	}

	if props.ID != "" {
		buttonArgs = append(buttonArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		buttonArgs = append(buttonArgs, attr)
	}
	buttonArgs = append(buttonArgs, args...)

	return html.Button(buttonArgs...)
}

// Next creates a next navigation button
func Next(props NextProps, args ...html.ButtonArg) html.Node {
	buttonArgs := []html.ButtonArg{
		html.AClass(classnames.Merge("absolute right-2 top-1/2 transform -translate-y-1/2 p-2 rounded-full bg-black/20 text-white hover:bg-black/40 focus:outline-none", props.Class)),
		html.AData("pui-carousel-next", ""),
		html.AAria("label", "Next slide"),
		html.AType("button"),
		lucide.ChevronRight(html.AClass("h-4 w-4")),
	}

	if props.ID != "" {
		buttonArgs = append(buttonArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		buttonArgs = append(buttonArgs, attr)
	}
	buttonArgs = append(buttonArgs, args...)

	return html.Button(buttonArgs...)
}

// Indicators creates carousel indicators for navigation
func Indicators(props IndicatorsProps, args ...html.DivArg) html.Node {
	divArgs := []html.DivArg{
		html.AClass(classnames.Merge("absolute bottom-4 left-1/2 transform -translate-x-1/2 flex gap-2", props.Class)),
	}

	if props.ID != "" {
		divArgs = append(divArgs, html.AId(props.ID))
	}
	for _, attr := range props.Attrs {
		divArgs = append(divArgs, attr)
	}

	// Add indicator buttons
	indicatorButtons := make([]html.DivArg, 0, props.Count)
	for i := 0; i < props.Count; i++ {
		buttonClass := "w-3 h-3 rounded-full bg-foreground/30 hover:bg-foreground/50 focus:outline-none transition-colors"
		if i == 0 {
			buttonClass = classnames.Merge(buttonClass, "bg-primary")
		}

		button := html.Button(
			html.AClass(buttonClass),
			html.AData("pui-carousel-indicator", strconv.Itoa(i)),
			html.AData("pui-carousel-active", func() string {
				if i == 0 {
					return "true"
				}
				return "false"
			}()),
			html.AAria("label", fmt.Sprintf("Go to slide %d", i+1)),
			html.AType("button"),
		)
		indicatorButtons = append(indicatorButtons, button)
	}

	divArgs = append(divArgs, indicatorButtons...)
	divArgs = append(divArgs, args...)

	return html.Div(divArgs...)
}

func randomID(prefix string) string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return prefix + "-id"
	}
	return prefix + "-" + hex.EncodeToString(buf)
}

const carouselJS = `(function() {
  'use strict';

  const autoplays = new Map();
  let dragState = null;

  // Click handling for navigation
  document.addEventListener('click', (e) => {
    const prevBtn = e.target.closest('[data-pui-carousel-prev]');
    if (prevBtn) {
      const carousel = prevBtn.closest('[data-pui-carousel]');
      if (carousel) navigate(carousel, -1);
      return;
    }

    const nextBtn = e.target.closest('[data-pui-carousel-next]');
    if (nextBtn) {
      const carousel = nextBtn.closest('[data-pui-carousel]');
      if (carousel) navigate(carousel, 1);
      return;
    }

    const indicator = e.target.closest('[data-pui-carousel-indicator]');
    if (indicator) {
      const carousel = indicator.closest('[data-pui-carousel]');
      const index = parseInt(indicator.dataset.tuiCarouselIndicator);
      if (carousel && !isNaN(index)) {
        updateCarousel(carousel, index);
      }
    }
  });

  // Drag/swipe handling
  function startDrag(e) {
    const track = e.target.closest('[data-pui-carousel-track]');
    if (!track) return;

    const carousel = track.closest('[data-pui-carousel]');
    if (!carousel) return;

    e.preventDefault();
    const clientX = e.touches ? e.touches[0].clientX : e.clientX;

    dragState = {
      carousel,
      track,
      startX: clientX,
      currentX: clientX,
      startTime: Date.now()
    };

    track.style.cursor = 'grabbing';
    track.style.transition = 'none';
    stopAutoplay(carousel);
  }

  function doDrag(e) {
    if (!dragState) return;

    const clientX = e.touches ? e.touches[0].clientX : e.clientX;
    dragState.currentX = clientX;

    const diff = clientX - dragState.startX;
    const currentIndex = parseInt(dragState.carousel.dataset.tuiCarouselCurrent || '0');
    const offset = -currentIndex * 100 + (diff / dragState.track.offsetWidth) * 100;

    dragState.track.style.transform = 'translateX(' + offset + '%)';
  }

  function endDrag(e) {
    if (!dragState) return;

    const { carousel, track, startX, startTime } = dragState;
    const clientX = e.changedTouches ? e.changedTouches[0].clientX : (e.clientX || dragState.currentX);

    track.style.cursor = '';
    track.style.transition = '';

    const diff = startX - clientX;
    const velocity = Math.abs(diff) / (Date.now() - startTime);

    if (Math.abs(diff) > 50 || velocity > 0.5) {
      navigate(carousel, diff > 0 ? 1 : -1);
    } else {
      const currentIndex = parseInt(carousel.dataset.tuiCarouselCurrent || '0');
      updateCarousel(carousel, currentIndex);
    }

    dragState = null;

    if (carousel.dataset.tuiCarouselAutoplay === 'true' && !carousel.matches(':hover')) {
      startAutoplay(carousel);
    }
  }

  document.addEventListener('mousedown', startDrag);
  document.addEventListener('mousemove', doDrag);
  document.addEventListener('mouseup', endDrag);
  document.addEventListener('mouseleave', (e) => {
    if (e.target === document.documentElement) endDrag(e);
  });

  document.addEventListener('touchstart', startDrag, { passive: false });
  document.addEventListener('touchmove', doDrag, { passive: false });
  document.addEventListener('touchend', endDrag, { passive: false });

  // Navigation logic
  function navigate(carousel, direction) {
    const current = parseInt(carousel.dataset.tuiCarouselCurrent || '0');
    const items = carousel.querySelectorAll('[data-pui-carousel-item]');
    const count = items.length;

    if (count === 0) return;

    let next = current + direction;

    if (carousel.dataset.tuiCarouselLoop === 'true') {
      next = ((next % count) + count) % count;
    } else {
      next = Math.max(0, Math.min(next, count - 1));
    }

    updateCarousel(carousel, next);
  }

  function updateCarousel(carousel, index) {
    const track = carousel.querySelector('[data-pui-carousel-track]');
    const indicators = carousel.querySelectorAll('[data-pui-carousel-indicator]');
    const prevBtn = carousel.querySelector('[data-pui-carousel-prev]');
    const nextBtn = carousel.querySelector('[data-pui-carousel-next]');
    const items = carousel.querySelectorAll('[data-pui-carousel-item]');
    const count = items.length;

    carousel.dataset.tuiCarouselCurrent = index;

    if (track) {
      track.style.transform = 'translateX(-' + (index * 100) + '%)';
    }

    indicators.forEach((ind, i) => {
      ind.dataset.tuiCarouselActive = (i === index) ? 'true' : 'false';
      ind.classList.toggle('bg-primary', i === index);
      ind.classList.toggle('bg-foreground/30', i !== index);
    });

    const isLoop = carousel.dataset.tuiCarouselLoop === 'true';

    if (prevBtn) {
      prevBtn.disabled = !isLoop && index === 0;
      prevBtn.classList.toggle('opacity-50', prevBtn.disabled);
    }

    if (nextBtn) {
      nextBtn.disabled = !isLoop && index === count - 1;
      nextBtn.classList.toggle('opacity-50', nextBtn.disabled);
    }
  }

  // Autoplay functionality
  function startAutoplay(carousel) {
    if (carousel.dataset.tuiCarouselAutoplay !== 'true') return;

    stopAutoplay(carousel);

    const interval = parseInt(carousel.dataset.tuiCarouselInterval || '5000');
    const id = setInterval(() => {
      if (!document.contains(carousel)) {
        stopAutoplay(carousel);
        return;
      }

      if (carousel.matches(':hover') || dragState?.carousel === carousel) {
        return;
      }

      navigate(carousel, 1);
    }, interval);

    autoplays.set(carousel, id);
  }

  function stopAutoplay(carousel) {
    const id = autoplays.get(carousel);
    if (id) {
      clearInterval(id);
      autoplays.delete(carousel);
    }
  }

  // Intersection Observer for visibility management
  const observedCarousels = new WeakSet();
  const carouselObserver = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      const carousel = entry.target;

      // Initialize display on first observation
      if (!carousel.hasAttribute('data-pui-carousel-initialized')) {
        carousel.setAttribute('data-pui-carousel-initialized', 'true');
        const index = parseInt(carousel.dataset.tuiCarouselCurrent || '0');
        updateCarousel(carousel, index);
      }

      // Handle autoplay if enabled
      if (carousel.dataset.tuiCarouselAutoplay === 'true') {
        if (entry.isIntersecting) {
          startAutoplay(carousel);
        } else {
          stopAutoplay(carousel);
        }
      }
    });
  });

  // Observe all carousels for visibility and initialization
  function observeCarousels() {
    document.querySelectorAll('[data-pui-carousel]').forEach(carousel => {
      if (!observedCarousels.has(carousel)) {
        observedCarousels.add(carousel);
        carouselObserver.observe(carousel);
      }
    });
  }

  // Start observing
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', observeCarousels);
  } else {
    observeCarousels();
  }

  // Watch for dynamically added carousels
  new MutationObserver(observeCarousels).observe(document.body, {
    childList: true,
    subtree: true
  });
})();`

'use strict';

const SELECTOR_ITEM = '.slider__item';
const SELECTOR_ITEMS = '.slider__items';
const SELECTOR_WRAPPER = '.slider__wrapper';
const SELECTOR_PREV = '.slider__control[data-slide="prev"]';
const SELECTOR_NEXT = '.slider__control[data-slide="next"]';
const SELECTOR_INDICATOR = '.slider__indicators>li';

const SLIDER_TRANSITION_OFF = 'slider_disable-transition';
const CLASS_CONTROL = 'slider__control';
const CLASS_CONTROL_HIDE = 'slider__control_hide';
const CLASS_ITEM_ACTIVE = 'slider__item_active';
const CLASS_INDICATOR_ACTIVE = 'active';

function hasTouchDevice() {
  return !!('ontouchstart' in window || navigator.maxTouchPoints);
}

function hasElementInVew($elem) {
  const rect = $elem.getBoundingClientRect();
  const windowHeight = window.innerHeight || document.documentElement.clientHeight;
  const windowWidth = window.innerWidth || document.documentElement.clientWidth;
  const vertInView = rect.top <= windowHeight && rect.top + rect.height >= 0;
  const horInView = rect.left <= windowWidth && rect.left + rect.width >= 0;
  return vertInView && horInView;
}

class SliderByItchief {
  // configuration of the slider
  _config = {};
  // slider properties
  _widthItem = 0;
  _widthWrapper = 0;
  _itemsInVisibleArea = 0;
  _transform = 0; // текущее значение трансформации
  _transformStep = 0; // значение шага трансформации
  _intervalId = null;
  // elements of slider
  _$root = null; // root element of slider (default ".slider__item")
  _$wrapper = null; // element with class ".slider__wrapper"
  _$items = null; // element with class ".slider__items"
  _$itemList = null; // elements with class ".slider__item"
  _$controlPrev = null; // element with class .slider__control[data-slide="prev"]
  _$controlNext = null; // element with class .slider__control[data-slide="next"]
  _$indicatorList = null; // индикаторы
  // min and min order
  _minOrder = 0;
  _maxOrder = 0;
  // items with min and max order
  _$itemByMinOrder = null;
  _$itemByMaxOrder = null;
  // min and max value of translate
  _minTranslate = 0;
  _maxTranslate = 0;
  // default slider direction
  _direction = 'next';
  // determines whether the position of item needs to be determined
  _updateItemPositionFlag = false;
  _activeItems = [];
  _isTouchDevice = hasTouchDevice();

  // constructor
  constructor($elem, config) {
    this._init($elem, config);
    this._addEventListener();
  }

  // initial setup of slider
  _init($root, config) {
    // elements of slider
    this._$root = $root;
    this._$itemList = $root.querySelectorAll(SELECTOR_ITEM);
    this._$items = $root.querySelector(SELECTOR_ITEMS);
    this._$wrapper = $root.querySelector(SELECTOR_WRAPPER);
    this._$controlPrev = $root.querySelector(SELECTOR_PREV);
    this._$controlNext = $root.querySelector(SELECTOR_NEXT);
    this._$indicatorList = $root.querySelectorAll(SELECTOR_INDICATOR);
    // create some constants
    const $itemList = this._$itemList;
    const widthItem = parseFloat(getComputedStyle($itemList[0]).width);
    const widthWrapper = parseFloat(getComputedStyle(this._$wrapper).width);
    const itemsInVisibleArea = Math.round(widthWrapper / widthItem);
    // initial setting properties
    this._widthItem = widthItem;
    this._widthWrapper = widthWrapper;
    this._itemsInVisibleArea = itemsInVisibleArea;
    this._transformStep = 100 / itemsInVisibleArea;
    this._config = config;
    // initial setting order and translate items
    $itemList.forEach(($item, position) => {
      $item.dataset.index = position;
      $item.dataset.order = position;
      $item.dataset.translate = 0;
      if (position < itemsInVisibleArea) {
        this._activeItems.push(position);
      }
    });
    this._updateClassForActiveItems();
    // hide prev arrow for non-infinite slider
    if (!this._config.loop) {
      if (this._$controlPrev) {
        this._$controlPrev.classList.add(CLASS_CONTROL_HIDE);
      }
      return;
    }
    // translate last item before first
    const count = $itemList.length - 1;
    const translate = -$itemList.length * 100;
    $itemList[count].dataset.order = -1;
    $itemList[count].dataset.translate = -$itemList.length * 100;
    $itemList[count].style.transform = `translateX(${translate}%)`;
    // update values of extreme properties
    this._updateExtremeProperties();
    this._updateIndicators();
    // calling _autoplay
    this._autoplay();
  }

  // подключения обработчиков событий для слайдера
  _addEventListener() {
    const $root = this._$root;

    // on click
    $root.addEventListener('click', this._eventHandler.bind(this));

    // on hover
    if (this._config.autoplay && this._config.pauseOnHover) {
      $root.addEventListener('mouseenter', () => {
        this._autoplay('stop');
      });
      $root.addEventListener('mouseleave', () => {
        this._autoplay();
      });
    }

    // on resize
    if (this._config.refresh) {
      window.addEventListener('resize', () => {
        window.requestAnimationFrame(this._refresh.bind(this));
      });
    }

    // on transitionstart and transitionend
    if (this._config.loop) {
      this._$items.addEventListener('transitionstart', () => {
        // transitionstart
        this._updateItemPositionFlag = true;
        window.requestAnimationFrame(this._updateItemPosition.bind(this));
      });
      this._$items.addEventListener('transitionend', () => {
        // transitionend
        this._updateItemPositionFlag = false;
      });
    }

    // on touchstart and touchend
    if (this._isTouchDevice && this._config.swipe) {
      $root.addEventListener('touchstart', e => {
        this._touchStartCoord = e.changedTouches[0].clientX;
      });
      $root.addEventListener('touchend', e => {
        const touchEndCoord = e.changedTouches[0].clientX;
        const delta = touchEndCoord - this._touchStartCoord;
        if (delta > 50) {
          this._moveToPrev();
        } else if (delta < -50) {
          this._moveToNext();
        }
      });
    }

    // on mousedown and mouseup
    if (!this._isTouchDevice && this._config.swipe) {
      $root.addEventListener('mousedown', e => {
        this._touchStartCoord = e.clientX;
      });
      $root.addEventListener('mouseup', e => {
        const touchEndCoord = e.clientX;
        const delta = touchEndCoord - this._touchStartCoord;
        if (delta > 50) {
          this._moveToPrev();
        } else if (delta < -50) {
          this._moveToNext();
        }
      });
    }
  }

  // update values of extreme properties
  _updateExtremeProperties() {
    const $itemList = this._$itemList;
    this._minOrder = +$itemList[0].dataset.order;
    this._maxOrder = this._minOrder;
    this._$itemByMinOrder = $itemList[0];
    this._$itemByMaxOrder = $itemList[0];
    this._minTranslate = +$itemList[0].dataset.translate;
    this._maxTranslate = this._minTranslate;
    $itemList.forEach($item => {
      const order = +$item.dataset.order;
      if (order < this._minOrder) {
        this._minOrder = order;
        this._$itemByMinOrder = $item;
        this._minTranslate = +$item.dataset.translate;
      } else if (order > this._maxOrder) {
        this._maxOrder = order;
        this._$itemByMaxOrder = $item;
        this._minTranslate = +$item.dataset.translate;
      }
    });
  }

  // update position of item
  _updateItemPosition() {
    if (!this._updateItemPositionFlag) {
      return;
    }
    const $wrapper = this._$wrapper;
    const $wrapperClientRect = $wrapper.getBoundingClientRect();
    const widthHalfItem = $wrapperClientRect.width / this._itemsInVisibleArea / 2;
    const count = this._$itemList.length;
    if (this._direction === 'next') {
      const wrapperLeft = $wrapperClientRect.left;
      const $min = this._$itemByMinOrder;
      let translate = this._minTranslate;
      const clientRect = $min.getBoundingClientRect();
      if (clientRect.right < wrapperLeft - widthHalfItem) {
        $min.dataset.order = this._minOrder + count;
        translate += count * 100;
        $min.dataset.translate = translate;
        $min.style.transform = `translateX(${translate}%)`;
        // update values of extreme properties
        this._updateExtremeProperties();
      }
    } else {
      const wrapperRight = $wrapperClientRect.right;
      const $max = this._$itemByMaxOrder;
      let translate = this._maxTranslate;
      const clientRect = $max.getBoundingClientRect();
      if (clientRect.left > wrapperRight + widthHalfItem) {
        $max.dataset.order = this._maxOrder - count;
        translate -= count * 100;
        $max.dataset.translate = translate;
        $max.style.transform = `translateX(${translate}%)`;
        // update values of extreme properties
        this._updateExtremeProperties();
      }
    }
    // updating...
    requestAnimationFrame(this._updateItemPosition.bind(this));
  }

  // _updateClassForActiveItems
  _updateClassForActiveItems() {
    const activeItems = this._activeItems;
    this._$itemList.forEach($item => {
      const index = +$item.dataset.index;
      if (activeItems.includes(index)) {
        $item.classList.add(CLASS_ITEM_ACTIVE);
      } else {
        $item.classList.remove(CLASS_ITEM_ACTIVE);
      }
    });
  }

  // _updateIndicators
  _updateIndicators() {
    const $indicatorList = this._$indicatorList;
    if (!$indicatorList.length) {
      return;
    }
    this._$itemList.forEach(($item, index) => {
      if ($item.classList.contains(CLASS_ITEM_ACTIVE)) {
        $indicatorList[index].classList.add(CLASS_INDICATOR_ACTIVE);
      } else {
        $indicatorList[index].classList.remove(CLASS_INDICATOR_ACTIVE);
      }
    });
  }

  // move slides
  _move() {
    if (!hasElementInVew(this._$root)) {
      return;
    }

    const step = this._direction === 'next' ? -this._transformStep : this._transformStep;
    const transform = this._transform + step;

    if (!this._config.loop) {
      const endTransformValue =
        this._transformStep * (this._$itemList.length - this._itemsInVisibleArea);
      if (transform < -endTransformValue || transform > 0) {
        return;
      }
      this._$controlPrev.classList.remove(CLASS_CONTROL_HIDE);
      this._$controlNext.classList.remove(CLASS_CONTROL_HIDE);
      if (transform === -endTransformValue) {
        this._$controlNext.classList.add(CLASS_CONTROL_HIDE);
      } else if (transform === 0) {
        this._$controlPrev.classList.add(CLASS_CONTROL_HIDE);
      }
    }

    const activeIndex = [];
    if (this._direction === 'next') {
      this._activeItems.forEach(index => {
        let newIndex = ++index;
        if (newIndex > this._$itemList.length - 1) {
          newIndex -= this._$itemList.length;
        }
        activeIndex.push(newIndex);
      });
    } else {
      this._activeItems.forEach(index => {
        let newIndex = --index;
        if (newIndex < 0) {
          newIndex += this._$itemList.length;
        }
        activeIndex.push(newIndex);
      });
    }
    this._activeItems = activeIndex;
    this._updateClassForActiveItems();
    this._updateIndicators();

    this._transform = transform;
    this._$items.style.transform = `translateX(${transform}%)`;
  }

  // _moveToNext
  _moveToNext() {
    this._direction = 'next';
    this._move();
  }

  // _moveToPrev
  _moveToPrev() {
    this._direction = 'prev';
    this._move();
  }

  // _moveTo
  _moveTo(index) {
    const $indicatorList = this._$indicatorList;
    let nearestIndex = null;
    let diff = null;
    $indicatorList.forEach($indicator => {
      if ($indicator.classList.contains(CLASS_INDICATOR_ACTIVE)) {
        const slideTo = +$indicator.dataset.slideTo;
        if (diff === null) {
          nearestIndex = slideTo;
          diff = Math.abs(index - nearestIndex);
        } else {
          if (Math.abs(index - slideTo) < diff) {
            nearestIndex = slideTo;
            diff = Math.abs(index - nearestIndex);
          }
        }
      }
    });
    diff = index - nearestIndex;
    if (diff === 0) {
      return;
    }
    this._direction = diff > 0 ? 'next' : 'prev';
    for (let i = 1; i <= Math.abs(diff); i++) {
      this._move();
    }
  }

  // обработчик click для слайдера
  _eventHandler(e) {
    const $target = e.target;
    this._autoplay('stop');
    if ($target.classList.contains(CLASS_CONTROL)) {
      // при клике на кнопки влево и вправо
      e.preventDefault();
      this._direction = $target.dataset.slide;
      this._move();
    } else if ($target.dataset.slideTo) {
      // при клике на индикаторы
      const index = +$target.dataset.slideTo;
      this._moveTo(index);
    }
    this._autoplay();
  }

  // _autoplay
  _autoplay(action) {
    if (!this._config.autoplay) {
      return;
    }
    if (action === 'stop') {
      clearInterval(this._intervalId);
      this._intervalId = null;
      return;
    }
    if (this._intervalId === null) {
      this._intervalId = setInterval(() => {
        this._direction = 'next';
        this._move();
      }, this._config.interval);
    }
  }

  // _refresh
  _refresh() {
    // create some constants
    const $itemList = this._$itemList;
    const widthItem = parseFloat(getComputedStyle($itemList[0]).width);
    const widthWrapper = parseFloat(getComputedStyle(this._$wrapper).width);
    const itemsInVisibleArea = Math.round(widthWrapper / widthItem);

    if (itemsInVisibleArea === this._itemsInVisibleArea) {
      return;
    }

    this._autoplay('stop');

    this._$items.classList.add(SLIDER_TRANSITION_OFF);
    this._$items.style.transform = 'translateX(0)';

    // setting properties after reset
    this._widthItem = widthItem;
    this._widthWrapper = widthWrapper;
    this._itemsInVisibleArea = itemsInVisibleArea;
    this._transform = 0;
    this._transformStep = 100 / itemsInVisibleArea;
    this._updateItemPositionFlag = false;
    this._activeItems = [];

    // setting order and translate items after reset
    $itemList.forEach(($item, position) => {
      $item.dataset.index = position;
      $item.dataset.order = position;
      $item.dataset.translate = 0;
      $item.style.transform = 'translateX(0)';
      if (position < itemsInVisibleArea) {
        this._activeItems.push(position);
      }
    });

    this._updateClassForActiveItems();

    window.requestAnimationFrame(() => {
      console.log(this);
      this._$items.classList.remove(SLIDER_TRANSITION_OFF);
    });

    // hide prev arrow for non-infinite slider
    if (!this._config.loop) {
      if (this._$controlPrev) {
        this._$controlPrev.classList.add(CLASS_CONTROL_HIDE);
      }
      return;
    }

    // translate last item before first
    const count = $itemList.length - 1;
    const translate = -$itemList.length * 100;
    $itemList[count].dataset.order = -1;
    $itemList[count].dataset.translate = -$itemList.length * 100;
    $itemList[count].style.transform = `translateX(${translate}%)`;
    // update values of extreme properties
    this._updateExtremeProperties();
    this._updateIndicators();
    // calling _autoplay
    this._autoplay();
  }

  // public
  next() {
    this._moveToNext();
  }
  prev() {
    this._moveToPrev();
  }
  moveTo(index) {
    this._moveTo(index);
  }
  refresh() {
    this._refresh();
  }
}

/*document.querySelectorAll('[data-slider="chiefslider"]').forEach($slider => {
  const dataset = $slider.dataset;
  const loop = dataset.loop ? true : false;
  const autoplay = dataset.autoplay ? true : false;
  const interval = dataset.interval ? parseInt(dataset.interval) : 5000;
  const pauseOnHover = dataset.pauseOnHover ? true : false;
  new SliderByItchief($element, {
    loop: loop,
    autoplay: autoplay,
    interval: interval,
    pauseOnHover: pauseOnHover,
  });
});*/

const $slider = document.querySelector('[data-slider="chiefslider"]');
const slider = new SliderByItchief($slider, {
  loop: true, // whether to enable infinity loop (default: true)
  autoplay: false, // whether to enable autoplay (default: false)
  interval: 500, // autoplay interval in milliseconds (default: 5000)
  pauseOnHover: true, // whether to stop autoplay while a slider is hovered
  refresh: true, // should the slider be updated when the viewport is resized
  swipe: true, // enable swiping
});

// perPage
// perMove
// pagination

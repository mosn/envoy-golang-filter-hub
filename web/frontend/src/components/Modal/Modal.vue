<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue';

const props = defineProps({
  /**
   * 是否可见
   */
  visible: {
    type: Boolean,
    default: false,
  },
  visibleModifiers: {
    type: Object,
    default: () => ({}),
  },
  /**
   * 确认按钮文本
   */
  okText: {
    type: String,
    default: 'Ok',
  },
  /**
   * 取消按钮文本
   */
  cancelText: {
    type: String,
    default: 'Cancel',
  },
  /**
   * 按下按钮后是否关闭对话框
   */
  autoClose: {
    type: Boolean,
    default: true,
  },
  /**
   * 对话框内容
   */
  content: {
    type: String,
    default: '',
  },
  /**
   * 对话框内容是否居中
   */
  centerContent: {
    type: Boolean,
    default: false,
  },
  /**
   * 对话框标题
   */
  title: {
    type: String,
    default: 'Alert',
  },
  /**
   * 是否隐藏取消按钮
   */
  hideCancel: {
    type: Boolean,
    default: false,
  },
  /**
   * 是否隐藏确认按钮
   */
  hideOk: {
    type: Boolean,
    default: false,
  },
  /**
   * 是否显示遮罩
   */
  mask: {
    type: Boolean,
    default: true,
  },
  /**
   * 是否渲染到body元素
   */
  renderToBody: {
    type: Boolean,
    default: true,
  },
  dynamic: {
    type: Boolean,
    default: false,
  },
  moveable: {
    type: Boolean,
    default: false,
  },
});

const modalEl = ref();

const emit = defineEmits(['cancel', 'ok', 'open', 'close', 'update:visible']);

const _visible = !props.visibleModifiers['model']
  ? ref(props.visible)
  : computed({
    get() {
      return props.visible;
    },
    set(value) {
      emit('update:visible', value);
    },
  });

const close = () => {
  _visible.value = false;
  if (props.dynamic) {
    unmountedHandler();
    modalEl.value.parentNode &&
      modalEl.value.parentNode.parentNode &&
      modalEl.value.parentNode.parentNode.removeChild(modalEl.value.parentNode);
  }
};

const open = () => {
  _visible.value = true;
};

const onCancelButtonClick = () => {
  if (!_visible.value) return;
  emit('cancel');
  if (props.autoClose) {
    close();
  }
};

const onOkButtonClick = () => {
  if (!_visible.value) return;
  emit('ok');
  if (props.autoClose) {
    close();
  }
};

const handleOpen = () => {
  if (_visible.value) {
    emit('open');
  }
};

const handleClose = () => {
  if (!_visible.value) {
    emit('close');
  }
};

const onEscKeyUp = (e: KeyboardEvent) => {
  if (e.keyCode === 27) {
    onCancelButtonClick();
  }
};

onMounted(() => {
  document.addEventListener('keyup', onEscKeyUp);
  document.addEventListener('pointermove', onMovePointerMove);
  document.addEventListener('pointerup', onMovePointerUp);
  addEventListener('resize', onWindowResize);
});

const unmountedHandler = () => {
  document.removeEventListener('keyup', onEscKeyUp);
  document.removeEventListener('pointermove', onMovePointerMove);
  document.removeEventListener('pointerup', onMovePointerUp);
  removeEventListener('resize', onWindowResize);
};

onUnmounted(unmountedHandler);

const modalContainerEl = ref();
let pointerId = -1;
const moveX = ref(0);
const moveY = ref(0);

const onWindowResize = () => {
  if (!props.moveable) return;
  const screenW = (window.innerWidth - modalContainerEl.value.offsetWidth) / 2;
  const screenH =
    (window.innerHeight - modalContainerEl.value.offsetHeight) / 2;

  if (moveX.value > screenW) {
    moveX.value = screenW;
  } else if (moveX.value < -screenW) {
    moveX.value = -screenW;
  }
  if (moveY.value > screenH) {
    moveY.value = screenH;
  } else if (moveY.value < -screenH) {
    moveY.value = -screenH;
  }
};

const onMovePointerDown = (e: PointerEvent) => {
  if (!props.moveable) return;
  if (pointerId !== -1) return;
  pointerId = e.pointerId;
};

const onMovePointerMove = (e: PointerEvent) => {
  if (!props.moveable) return;
  if (e.pointerId !== pointerId) return;
  moveX.value += e.movementX;
  moveY.value += e.movementY;
  const screenW = (window.innerWidth - modalContainerEl.value.offsetWidth) / 2;
  const screenH =
    (window.innerHeight - modalContainerEl.value.offsetHeight) / 2;

  if (moveX.value > screenW) {
    moveX.value = screenW;
  } else if (moveX.value < -screenW) {
    moveX.value = -screenW;
  }
  if (moveY.value > screenH) {
    moveY.value = screenH;
  } else if (moveY.value < -screenH) {
    moveY.value = -screenH;
  }
};

const onMovePointerUp = (e: PointerEvent) => {
  if (!props.moveable) return;
  if (pointerId !== e.pointerId) return;
  pointerId = -1;
};

defineExpose({ close, open });
</script>

<template>
  <teleport to="body" :disabled="!renderToBody">
    <transition name="fade" @after-enter="handleOpen" @after-leave="handleClose">
      <div v-if="_visible" ref="modalEl" class="modal">
        <div v-if="mask" class="mask" @click="onCancelButtonClick"></div>
        <div class="transform-container" :style="{ transform: `translate(${moveX}px,${moveY}px)` }">
          <div ref="modalContainerEl" class="modal-container" :class="{ moveable }">
            <div class="title" @pointerdown="onMovePointerDown">
              <slot name="header">
                <div class="title-default">{{ title }}</div>
              </slot>
            </div>
            <div class="content" :class="{ center: centerContent }">
              <slot>{{ content }}</slot>
            </div>
            <div v-if="!hideCancel || !hideOk" class="buttons">
              <div v-if="!hideCancel" class="cancel-btn" @click="onCancelButtonClick">
                {{ cancelText }}
              </div>
              <div v-if="!hideOk" class="ok-btn" @click="onOkButtonClick">
                {{ okText }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </teleport>
</template>

<style lang="scss" scoped>
.modal {
  @apply z-91 fixed;

  .mask {
    @apply fixed top-0 left-0 right-0 bottom-0 bg-black/20 dark-bg-black/80;
  }

  .transform-container {
    @apply fixed top-0 left-0 right-0 bottom-0 pointer-events-none;
  }

  .modal-container {
    @apply max-h-85vh pointer-events-auto;
    @apply flex flex-col;
    @apply border-1.5 border-primary/30;
    @apply fixed top-1/2 left-1/2 min-w-xs sm-min-w-sm md-max-w-3xl lg-max-w-4xl bg-white dark-bg-black rounded-xl p-5 transform-gpu -translate-x-1/2 -translate-y-1/2;

    .title {
      @apply pb-2 font-bold select-none;

      .title-default {
        @apply text-2xl text-primary text-center;
      }
    }

    &.moveable {
      .title {
        @apply cursor-move;
      }
    }

    .content {
      @apply text-lg text-primary/80 whitespace-pre-wrap;
      @apply overflow-auto;

      &.center {
        @apply text-center;
      }
    }

    .buttons {
      @apply flex justify-end space-x-4 mt-3;

      .cancel-btn,
      .ok-btn {
        @apply py-1.5 px-5 rounded-lg cursor-pointer motion-safe-transition select-none transform-gpu h-9;

        &:active {
          @apply scale-90;
        }
      }

      .cancel-btn {
        @apply bg-gray-400 bg-opacity-20 text-gray-600 motion-safe-transition opacity-100 pointer-events-auto;

        &:hover {
          @apply bg-opacity-30;
        }

        &.hide {
          @apply opacity-0 pointer-events-none;
        }
      }

      .ok-btn {
        @apply bg-primary text-white;
        @apply flex items-center;

        &:hover {
          @apply bg-opacity-90;
        }

        &.loading {
          @apply opacity-50 pointer-events-none;
          @apply max-w-15 text-2xl;
        }
      }
    }
  }
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.15s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>

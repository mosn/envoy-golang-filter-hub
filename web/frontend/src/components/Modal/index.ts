import { h, render, type VNode } from 'vue'
import { isDev } from '@/utils/env'
import Modal from './Modal.vue'

export type PromptModalCallbackFunc = (
  data: {
    modal: VNode
    methods: {
      close: () => void
      open: () => void
    }
  },
  inputResult: string
) => void

export interface PromptModalOptions {
  title?: string
  content?: string
  okText?: string
  cancelText?: string
  hideOk?: boolean
  hideCancel?: boolean
  placeholder?: string
  centerContent?: boolean
  autoClose?: boolean
  mask?: boolean
  onOk?: PromptModalCallbackFunc
  onCancel?: PromptModalCallbackFunc
}

export interface AlertModalOptions {
  title?: string
  content?: string
  okText?: string
  cancelText?: string
  hideOk?: boolean
  hideCancel?: boolean
  centerContent?: boolean
  mask?: boolean
  onOk?: () => void
  onCancel?: () => void
}

const createAlertModal = ({
  title = 'Alert',
  content = '',
  okText = 'Ok',
  cancelText = 'Cancel',
  hideOk = false,
  hideCancel = false,
  centerContent = false,
  mask = true,
  onOk = () => {},
  onCancel = () => {}
}: AlertModalOptions) => {
  const container = document.createElement('div')
  const vnode = h(Modal, {
    title,
    content,
    okText,
    cancelText,
    hideOk,
    hideCancel,
    centerContent,
    dynamic: true,
    visible: true,
    renderToBody: false,
    mask,
    onOk: () => {
      document.body.style.overflow = 'auto'
      onOk()
    },
    onCancel: () => {
      document.body.style.overflow = 'auto'
      onCancel()
    }
  })
  render(vnode, container)
  document.body.appendChild(container)
  document.body.style.overflow = 'hidden'
  return vnode
}

const createPromptModal = ({
  title = 'Prompt',
  content = '',
  okText = 'Ok',
  cancelText = 'Cancel',
  hideOk = false,
  hideCancel = false,
  placeholder = '',
  centerContent = false,
  autoClose = true,
  mask = true,
  onOk = () => {},
  onCancel = () => {}
}: PromptModalOptions) => {
  const container = document.createElement('div')
  let inputResult = ''
  const input = h('input', {
    placeholder,
    class:
      'bg-white dark-bg-gray-8 w-full px-2 py-1 rounded-xl outline-transparent transition border-gray-2 dark-border-gray-7 border-1.5',
    onInput: (e: InputEvent) => {
      inputResult = (e.target as HTMLInputElement).value
    }
  })
  const vnode = h(
    Modal,
    {
      title,
      content,
      okText,
      cancelText,
      hideOk,
      hideCancel,
      autoClose,
      centerContent,
      mask,
      dynamic: true,
      visible: true,
      renderToBody: false,
      onOk: () => {
        document.body.style.overflow = 'auto'
        onOk({ modal: vnode, methods: vnode.component!.exposed as any }, inputResult)
      },
      onCancel: () => {
        document.body.style.overflow = 'auto'
        onCancel({ modal: vnode, methods: vnode.component!.exposed as any }, inputResult)
      }
    },
    {
      default: () => [h('p', { class: 'text-primary' }, content), input]
    }
  )
  render(vnode, container)
  document.body.appendChild(container)
  document.body.style.overflow = 'hidden'
  return vnode
}

// export to window.debug if in debug mode
if (isDev) {
  window.debug = {
    ...(window.debug || {}),
    modal: {
      createAlertModal,
      createPromptModal
    }
  }
}

export { createAlertModal, createPromptModal }

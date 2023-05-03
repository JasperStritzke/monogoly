import {Directive, DirectiveBinding} from "vue";

interface ClickOutSideDirectiveBinding extends DirectiveBinding {
    value: () => void
}

interface HTMLClickOutsideElement extends HTMLElement {
    listener: (e: MouseEvent) => void
}

export const clickOutSide: Directive = {
    mounted(el: HTMLClickOutsideElement, binding: ClickOutSideDirectiveBinding) {
        el.listener = (e: MouseEvent) => {
            if (!(el == e.target || el.contains(e.target as HTMLElement))) {
                binding.value()
            }
        }
        setTimeout(() => document.addEventListener("click", el.listener), 0)
    },
    unmounted(el: HTMLClickOutsideElement) {
        document.removeEventListener("click", el.listener)
    }
}
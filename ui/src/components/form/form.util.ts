//extremely necessary form validation - could've been done completely without JS, only CSS and pseudo-classes... but here I go ^^ xd
import {inject, InjectionKey, onBeforeUnmount, onMounted, provide, ref, Ref} from "vue";

export type FieldValidationResult = boolean | string
export type FieldValidationRule = (value: any) => FieldValidationResult
export type FieldRules = Array<FieldValidationRule>
export type FormValidationMethod = () => boolean;
export type FormRegistrationMethod = (method: FormValidationMethod) => void;

export interface ModelInternalForm {
    lazyValidation?: Ref<boolean>,
    register: FormRegistrationMethod,
    unregister: FormRegistrationMethod,
}

const formSymbol = Symbol() as InjectionKey<ModelInternalForm>
//hooks
/**
 * hook to set up a local form, working in combination with {@link useField}
 * @param isLazyValidation
 * @return validate function
 */
function useUiForm(isLazyValidation: Ref<boolean> = ref(false)): () => boolean {
    const lazyValidation = ref(isLazyValidation)

    const validationMethods: Set<FormValidationMethod> = new Set<FormValidationMethod>()

    function register(method: FormValidationMethod) {
        validationMethods.add(method)
    }

    function unregister(method: FormValidationMethod) {
        validationMethods.delete(method)
    }

    function validate(): boolean {
        let valid = true

        for (let method of validationMethods) {
            if (!method()) {
                valid = false
            }
        }

        return valid
    }

    provide(
        formSymbol,
        {
            lazyValidation,
            register,
            unregister
        }
    )

    return validate
}

export {useUiForm}

export interface ModelFieldHandle {
    validate: () => boolean,
    resetValidation: () => void,
    error: Ref<string>,
    isLazyValidation?: Ref<boolean>,
}

function useField(modelValueRef: Ref, rules: Ref<FieldRules>): ModelFieldHandle {
    let uiForm = inject(formSymbol, undefined)

    const error = ref<string>("")

    function validate(): boolean {
        error.value = ""

        for (let rule of rules.value) {
            let validationResult = rule(modelValueRef.value)

            if (validationResult === true) {
                continue
            }

            if (typeof validationResult !== "string" || validationResult.length === 0) {
                validationResult = "Invalid field"
            }

            error.value = validationResult
            return false
        }

        return true
    }

    //if not, field is in standalone mode
    if (uiForm) {
        onMounted(() => {
            uiForm!.register(validate)
        })

        onBeforeUnmount(() => {
            uiForm!.unregister(validate)
        })
    }

    return {
        validate,
        error,
        resetValidation: () => error.value = "",
        isLazyValidation: uiForm?.lazyValidation
    }
}

export {useField}

export const RequiredRule: FieldValidationRule = (v: any) => !!v || "Field is required"
export const MinMaxRule = (min: number = 0, max: number = Number.MAX_VALUE): FieldValidationRule => {
    return (v: any) => {
        if(+v < min) return `Can't be less than ${min}`
        if(+v > max) return `Can't be more than ${max}`
        return true
    }
}
export const MinMaxLengthRule = (min: number = 0, max: number = Number.MAX_VALUE): FieldValidationRule => {
    return (v: any) => {
        if(v.length < min) return `Can't be shorter than ${min}`
        if(v.length > max) return `Can't be longer than ${max}`
        return true
    }
}

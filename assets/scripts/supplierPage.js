function afterItemAdd(modalSelector) {
    const hasFlashError = document.querySelector('.feedback-error') != null
    closeModal(modalSelector, "input[name=product]")
    if (!hasFlashError) {
        alterSupplierItemCounter(".selected", ".s-count")
    }
}

function afterRemoveItem() {
    console.log("func triggered")
    const hasFlashError = document.querySelector('.feedback-error') != null
    if (!hasFlashError) {
        alterSupplierItemCounter(".selected", ".s-count", false)
    }
}

function handleItemListAfterRequest(event) {
    const path = event?.detail?.pathInfo?.finalRequestPath || ''
    if (path.includes('/app/suppliers/delete/product')) {
        afterRemoveItem()
    }
}

function alterSupplierItemCounter(parentClass, targetClass, isAdditive = true) {
        if (document.querySelector(parentClass)) {
            const target = document.querySelector(parentClass).querySelector(targetClass)
            let text = target.innerHTML
            const splitted = text.split(" ")
            isAdditive ? splitted[0] = parseInt(splitted[0]) + 1 : splitted[0] = parseInt(splitted[0]) - 1
            target.innerHTML = splitted.join(" ")
        }
}

function closeModal(modalSelector, inputselector) {
    const modal = document.getElementById(modalSelector)
    const input = modal.querySelector(inputselector)
    if (modal && input) {
        modal.close()
        input.value = ""
    }
}

function setupSupplierInformation() {
    const editBtn = document.querySelector('#edit-btn')
    const saveBtn = document.querySelector('#save-btn')
    const cancelBtn = document.querySelector('#cancel-btn')
    const viewMode = document.querySelector('#view-mode')
    const editMode = document.querySelector('#edit-mode')

    if (!editBtn || !saveBtn || !cancelBtn || !viewMode || !editMode) {
        return
    }

    editBtn.addEventListener('click', () => {
        viewMode.classList.add('is-hidden')
        editMode.classList.remove('is-hidden')
        editBtn.classList.add('is-hidden')
        saveBtn.classList.remove('is-hidden')
        cancelBtn.classList.remove('is-hidden')
    })

    ;[saveBtn, cancelBtn].forEach(btn => {
        btn.addEventListener('click', () => {
            viewMode.classList.remove('is-hidden')
            editMode.classList.add('is-hidden')
            editBtn.classList.remove('is-hidden')
            saveBtn.classList.add('is-hidden')
            cancelBtn.classList.add('is-hidden')
        })
    })
}

document.addEventListener('DOMContentLoaded', () => {
    setupSupplierInformation()
})

document.body.addEventListener('htmx:afterSwap', event => {
    const target = event?.detail?.target
    if (target.id === 'supplier-information-slot' || target.id === 'supplier-content-slot') {
        setupSupplierInformation(target)
    }
})
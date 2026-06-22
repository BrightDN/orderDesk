function initializeListeners() {
    const footer = document.querySelector(".order-footer")
    if (!footer) {
        console.info("The order footer is not defined")
        return
    }

    const btn = footer.querySelector("button")
    if (!btn) {
        console.info("Order footer button was not located correctly")
        return
    }

    btn.addEventListener("click", sendOrder)
}

async function sendOrder() {
    const rows = document.querySelectorAll(".order-table tbody tr")
    if (!rows || rows.length == 0) return loadFeedbackBanner()

    const data = Array.from(rows) 
    if (!data || data.length == 0) return loadFeedbackBanner()

    let supplierID = Number(document.querySelector("[data-supplierID]").getAttribute("data-supplierID"))
    if (!supplierID) return loadFeedbackBanner()
    
    const filteredOrderData = filterByQty(data)
    if (filteredOrderData.length == 0) return loadFeedbackBanner()

    const JSONData = JSON.stringify({
        "supplierID": supplierID,
        "items": generateOrderData(filteredOrderData)
    })

    const response = await fetch("/app/order/send", {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body: JSONData
})

if (!response.ok) {
    return loadFeedbackBanner()
}

    const html = await response.text()
    loadFeedbackBanner(html)
}

function filterByQty(data) {
    return data.filter( (data) => data.querySelector(".qty-input").value > 0)
}

function generateOrderData(data) {
    const items = []
    data.forEach((data) => {
        obj = { productName: data.querySelector(".order-product").textContent.trim(), qty: Number(data.querySelector(".qty-input").value) }
        items.push(obj)
    })
    return items
}

function loadFeedbackBanner(html) {
    if (!html) {
        if (typeof showFeedbackBanner === 'function') {
            showFeedbackBanner('error', 'Unable to send order. Please try again.');
            return
        }
        return
    }

    if (typeof insertFeedbackHtml === 'function') {
        insertFeedbackHtml(html)
        return
    }

    // Fallback: append raw HTML if helper is unavailable
    const wrapper = document.createElement('div')
    wrapper.innerHTML = html.trim()
    const banner = wrapper.querySelector('[data-feedback-banner]')
    if (banner) {
        document.body.insertBefore(banner, document.querySelector('.app') || document.body.firstChild)
        if (typeof setupFeedbackBanners === 'function') {
            setupFeedbackBanners()
        }
    }
}

initializeListeners()
document.body.addEventListener('htmx:afterSwap', event => {
    const target = event?.detail?.target
    if (target && target.id === 'item-list-slot') {
        initializeListeners()
    }
})

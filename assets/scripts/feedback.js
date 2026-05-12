function dismissFeedbackBanner(banner) {
  banner.classList.add('is-leaving');
  window.setTimeout(() => banner.remove(), 180);
}

function setupFeedbackBanners() {
  document.querySelectorAll('[data-feedback-banner]').forEach((banner) => {
    if (banner.dataset.feedbackReady === 'true') return;
    banner.dataset.feedbackReady = 'true';

    const close = banner.querySelector('[data-feedback-close]');
    const timeout = window.setTimeout(() => dismissFeedbackBanner(banner), 5000);

    close?.addEventListener('click', () => {
      window.clearTimeout(timeout);
      dismissFeedbackBanner(banner);
    });
  });
}

document.addEventListener('DOMContentLoaded', setupFeedbackBanners);
document.addEventListener('htmx:afterSwap', setupFeedbackBanners);

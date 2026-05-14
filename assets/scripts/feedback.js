function dismissFeedbackBanner(banner) {
  banner.classList.add('is-leaving');
  window.setTimeout(() => banner.remove(), 180);
}

function buildFeedbackBanner(type, message) {
  const banner = document.createElement('aside');
  banner.className = `feedback-banner feedback-${type}`;
  banner.setAttribute('role', 'status');
  banner.setAttribute('aria-live', 'polite');
  banner.dataset.feedbackBanner = '';

  const text = document.createElement('p');
  text.textContent = message;

  const close = document.createElement('button');
  close.className = 'feedback-close';
  close.type = 'button';
  close.setAttribute('aria-label', 'Close message');
  close.dataset.feedbackClose = '';
  close.innerHTML = '<svg viewBox="0 0 14 14" aria-hidden="true"><path d="M3 3l8 8M11 3l-8 8"/></svg>';

  banner.append(text, close);
  return banner;
}

function showFeedbackBanner(type, message) {
  if (!type || !message) return;

  const banner = buildFeedbackBanner(type, message);
  document.body.insertBefore(banner, document.querySelector('.app') || document.body.firstChild);
  setupFeedbackBanners();
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
document.addEventListener('feedback', (event) => {
  showFeedbackBanner(event.detail?.type, event.detail?.message);
});

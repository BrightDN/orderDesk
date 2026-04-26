if (document.querySelectorAll('.supplier-card')) {
  document.querySelectorAll('.supplier-card').forEach(card => {
    card.addEventListener('click', () => {
      document.querySelectorAll('.supplier-card').forEach(c => c.classList.remove('selected'));
      card.classList.add('selected');
    });
  });
}

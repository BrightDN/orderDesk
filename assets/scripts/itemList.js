document.querySelectorAll('.del-btn').forEach(btn => {
    btn.addEventListener('click', () => {
      btn.closest('.order-row').remove();
    });
  });
function modalOpener(select, selectValue, modalCssSelector){
let modal = document.querySelector(modalCssSelector);
  if(select.value == selectValue){
    modal.showModal();
  }
}
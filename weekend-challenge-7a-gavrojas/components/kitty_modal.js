import Store from '../services/Store.js';

export default class KittyModal extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: 'open' });

    fetch("components/kitty_modal.css")
    .then((res) => res.text())
    .then((text) => {
      const style = document.createElement("style")
      style.textContent = text;
      this.root.appendChild(style)
    })
  }

  connectedCallback() {
    const template = document.getElementById('kitty-modal-template').content.cloneNode(true);
    this.root.appendChild(template);

    window.addEventListener('kitty-selected', () => {
      const modal = this.root.querySelector('.modal');
      const selectedKitty = Store.selectedKitty;
      modal.querySelector('h2').textContent = selectedKitty.name;
      modal.querySelector('img').src = selectedKitty.thumbnail;
      modal.querySelector('img').alt = selectedKitty.name;
      modal.querySelector('p').textContent = `$${selectedKitty.price}`;
      modal.classList.add('active');
    });

    window.addEventListener('kitty-clear', () => {
      const modal = this.root.querySelector('.modal');
      modal.classList.remove('active');
    });

    const closeBtn = this.root.getElementById('modal-close')

    closeBtn.addEventListener('click', () => {
      window.dispatchEvent(new Event("kitty-clear"));
    });
  }
}

customElements.define('kitty-modal', KittyModal);
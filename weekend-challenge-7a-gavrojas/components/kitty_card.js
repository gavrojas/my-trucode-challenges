import Store from '../services/Store.js';

export default class KittyCard extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: 'open' });
    fetch("components/kitty_card.css")
    .then((res) => res.text())
    .then((text) => {
      const style = document.createElement("style")
      style.textContent = text;
      this.root.appendChild(style)
    })
  }

  connectedCallback() {
    const template = document.getElementById('kitty-card-template').content.cloneNode(true);
    const id = this.dataset.id;
    const kitty = Store.kitties.find(k => k.id === id);
    if (kitty) {
      template.querySelector('h2').textContent = kitty.name;
      template.querySelector('img').src = kitty.thumbnail;
      template.querySelector('img').alt = kitty.name;
      template.querySelector('p').textContent = `$${kitty.price}`;
    }
    this.root.appendChild(template);

    this.root.querySelector('article').addEventListener('click', () => {
      Store.selectedKitty = kitty;
    });
  }
}

customElements.define('kitty-card', KittyCard);

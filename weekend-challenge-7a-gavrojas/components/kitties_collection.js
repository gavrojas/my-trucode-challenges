import Store from '../services/Store.js';
export default class KittyCollection extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: 'open' });

    fetch("components/kitties_collection.css")
    .then((res) => res.text())
    .then((text) => {
      const style = document.createElement("style")
      style.textContent = text;
      this.root.appendChild(style)
    })
  }

  connectedCallback() {
    const template = document.getElementById('kitties-collection-template')
    const content = template.content.cloneNode(true);
    this.root.appendChild(content);

    const list = this.root.querySelector('.list');
    Store.kitties.forEach(kitty => {
      const card = document.createElement('kitty-card');
      card.dataset.id = kitty.id;
      list.appendChild(card);
    });
  }
}

customElements.define('kitty-collection', KittyCollection);

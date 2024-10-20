import Alpine from 'alpinejs';

declare global {
  interface Window {
    hello: unknown;
    Alpine: typeof Alpine;
  }
}

window.hello = function (name: string) {
  alert("Hello " + name);
};

document.addEventListener('alpine:init', () => {
  console.log("ok"); // TODO: init event
})

Alpine.start()



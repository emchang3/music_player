declare const Vue: any;

var todoFetch = (context, url, purpose) => {
  fetch(url, { method: 'get' })
    .then((response) => response.json())
    .then((resJson) => { if (purpose === 'read') context.items = resJson; });
}

var app = new Vue({
  el: '#app',
  created: function() {
    todoFetch(this, '/init', 'read');
  },
  data: { items: [], playing: false },
  delimiters: [ '[', ']' ],
  methods: {
    cont: function() {
      this.playing = true;
      todoFetch(this, '/cont', 'op');
    },
    goUp: function() {
      todoFetch(this, '/up', 'read');
    },
    init: function() {
      todoFetch(this, '/init', 'read');
    },
    pause: function() {
      this.playing = false;
      todoFetch(this, '/pause', 'op');
    },
    play: function() {
      this.playing = true;
      todoFetch(this, `/play?item=${name}`, 'op');
    },
    read: function(name) {
      todoFetch(this, `/ls?item=${name}`, 'read');
    },
    next: function() {
      todoFetch(this, '/next', 'op');
    }
  }
});
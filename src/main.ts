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
  data: { items: [] },
  delimiters: [ '[', ']' ],
  methods: {
    goUp: function() {
      todoFetch(this, '/up', 'read');
    },
    init: function() {
      todoFetch(this, '/init', 'read');
    },
    play: function(name) {
      console.log(name);
      todoFetch(this, `/play?item=${name}`, 'op');
    },
    playDir: function() {
      todoFetch(this, `/playdir`, 'op');
    },
    read: function(name) {
      todoFetch(this, `/ls?item=${name}`, 'read');
    },
    stop: function() {
      todoFetch(this, '/stop', 'op');
    },
    next: function() {
      todoFetch(this, '/next', 'op')
    }
  }
});
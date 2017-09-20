declare const Vue: any;

var todoFetch = (context, url, purpose) => {
  fetch(url, { method: 'get' })
    .then((response) => response.json())
    .then((resJson) => { if (purpose === 'read') context.items = resJson; });
}

var app = new Vue({
  el: '#app',
  created: function() {
    todoFetch(this, 'http://localhost:2001/init', 'read');
  },
  data: { items: [], playing: false },
  delimiters: [ '[', ']' ],
  methods: {
    goUp: function() {
      todoFetch(this, 'http://localhost:2001/up', 'read');
    },
    init: function() {
      todoFetch(this, 'http://localhost:2001/init', 'read');
    },
    play: function(name) {
      console.log(name);
      todoFetch(this, `http://localhost:2001/play?item=${name}`, 'op');
      this.playing = true;
    },
    playDir: function() {
      todoFetch(this, `http://localhost:2001/playdir`, 'op');
      this.playing = true;
    },
    read: function(name) {
      todoFetch(this, `http://localhost:2001/ls?item=${name}`, 'read');
    },
    stop: function() {
      todoFetch(this, 'http://localhost:2001/stop', 'op');
      this.playing = false;
    },
    next: function() {
      todoFetch(this, 'http://localhost:2001/next', 'op');
      this.playing = true;
    },
    stopOrPlayCurrent: function() {
      if (this.playing) {
        todoFetch(this, 'http://localhost:2001/stop', 'op');
        this.playing = false;
      } else {
        todoFetch(this, `http://localhost:2001/playdir`, 'op');
        this.playing = true;
      }
    }
  }
});
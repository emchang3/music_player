declare const Vue: any;

var todoFetch = (context, url) => {
  fetch(url, { method: 'get' })
    .then((response) => response.json())
    .then((resJson) => { console.log(context); context.items = resJson; });
}

var app = new Vue({
  el: '#app',
  created: function() {
    todoFetch(this, '/init');
  },
  data: { items: [] },
  delimiters: [ '[', ']' ],
  methods: {
    goUp: function() {
      todoFetch(this, '/up');
    },
    init: function() {
      todoFetch(this, '/init');
    },
    read: function(name) {
      todoFetch(this, `/ls?item=${name}`);
    }
  }
});
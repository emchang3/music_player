var todoFetch=function(t,o,h){fetch(o,{method:"get"}).then(function(t){return t.json()}).then(function(o){"read"===h&&(t.items=o)})},app=new Vue({el:"#app",created:function(){todoFetch(this,"http://localhost:2001/init","read")},data:{items:[],playing:!1},delimiters:["[","]"],methods:{goUp:function(){todoFetch(this,"http://localhost:2001/up","read")},init:function(){todoFetch(this,"http://localhost:2001/init","read")},play:function(t){console.log(t),todoFetch(this,"http://localhost:2001/play?item="+t,"op"),this.playing=!0},playDir:function(){todoFetch(this,"http://localhost:2001/playdir","op"),this.playing=!0},read:function(t){todoFetch(this,"http://localhost:2001/ls?item="+t,"read")},stop:function(){todoFetch(this,"http://localhost:2001/stop","op"),this.playing=!1},next:function(){todoFetch(this,"http://localhost:2001/next","op"),this.playing=!0},stopOrPlayCurrent:function(){this.playing?(todoFetch(this,"http://localhost:2001/stop","op"),this.playing=!1):(todoFetch(this,"http://localhost:2001/playdir","op"),this.playing=!0)}}});
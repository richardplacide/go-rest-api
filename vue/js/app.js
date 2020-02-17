var app = new Vue({
    el: '#vueapp',
    data: {
      message: 'Hello from Vue!',
      questions: []
    },
    methods: {
      fetchData() {
        fetch('http://localhost:8000/articles')
        .then( response => response.json() )
        .then( jsonDATA => this.questions = jsonDATA.results) 
        }
      }
    })
    
var app = new Vue({
    el: '#vueapp',
    data: {
      message: 'Hello from Vue!',
      articles: [],
    },
    created () {
      this.fetchData() 
    },
    methods: {
      fetchData() {
        fetch('http://localhost:8000/articles')
        .then( response => response.json() )
        //.then(json => console.log(json))
        .then(json => {
          this.articles = json;
          
        })
        .catch( err => console.log('Error'))
        
        }
      }
    })

    
    
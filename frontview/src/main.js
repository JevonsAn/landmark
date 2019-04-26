import Vue from 'vue'
import App from './App.vue'
import router from './router'
import iView from 'iview';
import axios from 'axios';
import 'iview/dist/styles/iview.css';

Vue.use(iView);
// Vue.use(axios);

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
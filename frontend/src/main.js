import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import i18n from "./i18n";
import "@/plugins/element.js";
import "@/plugins/axios.js"
import "@/plugins/meta.js"
import "@/plugins/selfAdaptive.js"

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  i18n,
  render: h => h(App)
}).$mount("#app");

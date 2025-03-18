import { createStore } from 'vuex';

export default createStore({
  state: {
    token: sessionStorage.getItem('token') || ''
  },
  mutations: {
    setToken(state, token) {
      state.token = token;
      sessionStorage.setItem('token', token);
    },
    logout(state) {
      state.token = '';
      sessionStorage.removeItem('token');
    }
  },
});

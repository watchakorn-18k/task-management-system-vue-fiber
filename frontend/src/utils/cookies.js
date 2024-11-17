import Cookies from 'js-cookie';
import { reactive } from "vue";

const stateLogin = reactive({
    isLogin: false,
    jwt: ''
});

const setCookieUserJWT = (jwt) => {
    const now = new Date();
    const expireTime = new Date(now.getTime() + 6 * 60 * 60 * 1000);
    Cookies.set('jwt', jwt, { expires: expireTime });
    stateLogin.isLogin = true;
    stateLogin.jwt = jwt;
};

const getCookieUserJWT = () => {
    const jwt = Cookies.get('jwt');
    if (jwt) {
        return jwt;
    }
    return null;
};

const removeCookieUserJWT = () => {
    Cookies.remove('jwt');
    stateLogin.isLogin = false;
    stateLogin.jwt = '';
};



export { setCookieUserJWT, getCookieUserJWT, removeCookieUserJWT, stateLogin };
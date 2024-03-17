import Cookies from 'js-cookie';
import { navigate } from 'svelte-routing';

export const isAuthenticated = () => {
	const authenticated = Cookies.get('PodiDosa');

	return authenticated ? true : false;
};

export const signOut = () => {
	Cookies.remove('PodiDosa');
	navigate('/');
};

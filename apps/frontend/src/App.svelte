<script lang="ts">
	import { Router, Route } from 'svelte-routing';
	import { fade } from 'svelte/transition';

	import HomePage from './routes/home.svelte';
	import LoginPage from './routes/login.svelte';
	import RegisterPage from './routes/register.svelte';
	import ProtectedPage from './routes/protected.svelte';

	const basepath = '/';

	const viewtransition = () => ({
		fn: fade,
		duration: 500,
	});
</script>

<Router {basepath}>
	<Route path="/" component="{HomePage}" />
	<Route path="/login" component="{LoginPage}" />
	<Route path="/register" component="{RegisterPage}" />

	<Route path="/protected" component="{ProtectedPage}" />
</Router>

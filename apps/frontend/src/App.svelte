<script lang="ts">
	import { Router, Route } from 'svelte-routing';
	// import { fade } from 'svelte/transition'

	import HomePage from './routes/home.svelte';
	import LoginPage from './routes/login.svelte';
	import RegisterPage from './routes/register.svelte';

	const basepath = '/';
</script>

<Router {basepath}>
	<Route path="/" component="{HomePage}" />
	<Route path="/login" component="{LoginPage}" />
	<Route Path="/register" component="{RegisterPage}" />
</Router>

<script lang="ts">
	import { Router, Route } from 'svelte-routing';
	// import { fade } from 'svelte/transition'

	import HomePage from './routes/home.svelte';
	import LoginPage from './routes/login.svelte';
	import RegisterPage from './routes/register.svelte';
	import Dashboard from './routes/dashboard.svelte';
	import Drafts from './routes/drafts.svelte'; 
	import Chats from './routes/chats.svelte'; 
	import Explore from './routes/explore.svelte'; 
	import WebsiteBuilder from './routes/websiteBuilder.svelte'; 
	import Payments from './routes/payments.svelte'; 
	import Contracts from './routes/contracts.svelte'; 
	import Profile from './routes/profile.svelte'; 
	import Settings from './routes/settings.svelte'; 

	const basepath = '/';
</script>

<Router {basepath}>
	<Route path="/" component="{HomePage}" />
	<Route path="/login" component="{LoginPage}" />
	<Route Path="/register" component="{RegisterPage}" />
	<Route path="/dashboard" component="{Dashboard}" />
	<Route path="/drafts" component="{Drafts}" />
	<Route path="/chats" component = "{Chats}" />
	<Route path="/explore" component="{Explore}" />
	<Route path="/websiteBuilder" component="{WebsiteBuilder}" />
	<Route path="/payments" component="{Payments}" />
	<Route path="/contracts" component="{Contracts}" />
	<Route path="/profile" component="{Profile}" />
	<Route path="/settings" component="{Settings}" />
</Router>

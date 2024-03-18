<script lang="ts">
	import { Router, Route } from 'svelte-routing';
	import { fade } from 'svelte/transition';

	import HomePage from './routes/home.svelte';
	import LoginPage from './routes/login.svelte';
	import RegisterPage from './routes/register.svelte';

	import DashboardPage from './routes/dashboard.svelte';
	import DraftsPage from './routes/drafts.svelte';
	import ExplorePage from './routes/explore.svelte';
	import WebsiteBuilderPage from './routes/websiteBuilder.svelte';
	import PaymentsPage from './routes/payments.svelte';
	import ProfilePage from './routes/profile.svelte';
	import SettingsPage from './routes/settings.svelte';
	import ProtectedPage from './routes/protected.svelte';
	import ChatPage from './routes/chat.svelte';
	import ContractsPage from './routes/contracts.svelte';
	import MarketplacePage from './routes/marketplace.svelte';

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

	<Route path="/dashboard" component="{DashboardPage}" />
	<Route path="/drafts" component="{DraftsPage}" />
	<Route path="/explore" component="{ExplorePage}" />
	<Route path="/websiteBuilder" component="{WebsiteBuilderPage}" />
	<Route path="/payments" component="{PaymentsPage}" />
	<Route path="/profile" component="{ProfilePage}" />
	<Route path="/settings" component="{SettingsPage}" />
	<Route path="/protected" component="{ProtectedPage}" />
	<Route path="/chat" component="{ChatPage}" />
	<Route path="/contracts" component="{ContractsPage}" />
	<Route path="/marketplace" component="{MarketplacePage}" />
</Router>

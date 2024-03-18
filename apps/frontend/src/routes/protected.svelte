<script lang="ts">
	import { navigate } from 'svelte-routing';
	import { isAuthenticated, signOut } from '../lib/middleware/auth';
	import { onMount } from 'svelte';

	onMount(() => {
		if (!isAuthenticated()) {
			navigate('/');
		}
	});
</script>

<h1 class="text-white">This page is protected</h1>

<button
	on:click="{signOut}"
	class="bg-partyPurple text-spaceGrey rounded py-2 font-bold uppercase"
>
	Sign Out
</button>

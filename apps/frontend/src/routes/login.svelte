<script lang="ts">
	import { Link, navigate } from 'svelte-routing';
	import Cookies from 'js-cookie';

	import UnmuteLogo from '../assets/logo.svg';
	import BrandGoogle from '../assets/icons/brand-google.svg';

	let errorText: string | undefined;

	type TCredentials = {
		email?: string;
		password?: string;
	};

	const onSumbit = (e: SubmitEvent) => {
		errorText = undefined;

		const formData = new FormData(
			(e.target as HTMLFormElement) ?? undefined
		);

		const credentials: TCredentials = {};

		for (let [key, value] of formData) {
			credentials[key.toString() as keyof TCredentials] =
				value.toString();
		}

		console.log(credentials);

		fetch('http://localhost:8080/api/login', {
			method: 'POST',
			cache: 'no-cache',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(credentials),
		})
			.then((res) => {
				return res.json();
			})
			.then((data) => {
				if (data.error) {
					errorText = data.error;
					return;
				}

				Cookies.set('PodiDosa', data.token);
				navigate('/protected');
			})
			.catch((err) => console.error(err));
	};
</script>

<div
	class="w-full min-h-screen bg-spaceGrey flex flex-col items-center justify-center gap-8 p-8"
>
	<img src="{UnmuteLogo}" alt="Unmute Logo" class="absolute top-8 left-8" />
	<div class="w-full sm:w-1/2 lg:w-1/3 text-white">
		<span class="uppercase text-partyPurple">Login</span><span
			class="border-l mx-3"
		></span><Link to="/register" class="uppercase text-darkWhite"
			>Sign Up</Link
		>
	</div>

	<div class="flex flex-col gap-2 w-full sm:w-1/2 lg:w-1/3">
		<h1 class="text-white text-2xl font-ownersTrialWide tracking-wider">
			Welcome,
		</h1>
		<p class="text-darkWhite">Please enter your details below to log in</p>
	</div>

	<form
		on:submit|preventDefault="{onSumbit}"
		class="flex flex-col w-full sm:w-1/2 lg:w-1/3"
	>
		<label
			for="username"
			class="text-white text-sm mb-2 uppercase font-semibold"
		>
			Username
		</label>
		<input
			required
			type="text"
			name="username"
			id="username"
			placeholder="johndoe"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2 text-white"
		/>

		<label
			for="password"
			class="text-white text-sm mt-8 mb-2 uppercase font-semibold"
		>
			Password
		</label>
		<input
			required
			type="password"
			name="password"
			id="password"
			placeholder="******"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2 text-white"
		/>

		<Link
			to="/forgot-password"
			class="text-xs text-partyPurple my-4 uppercase font-semibold"
		>
			Forgot Password?
		</Link>

		<button
			type="submit"
			class="bg-partyPurple text-spaceGrey rounded py-2 font-bold uppercase"
		>
			Login
		</button>
	</form>

	<button
		class="flex justify-center items-center gap-2 border border-partyPurple text-darkWhite bg-transparent hover:bg-partyPurple w-full sm:w-1/2 lg:w-1/3 rounded py-2 font-bold uppercase"
	>
		<img src="{BrandGoogle}" alt="Google Icon" class="h-4" />
		<span> Sign in with Google </span>
	</button>

	{#if errorText}
		<p class="text-white">{errorText}</p>
	{/if}
</div>

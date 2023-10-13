<script>
	import '../styles.css';
</script>

<nav class="flex items-center py-5">
	<img src="/starfield_logo.svg" class="h-16 w-32 mx-5 object-contain" alt="Starfield logo" />
	<ul class="justify-end w-full flex space-x-5 mx-10">
		<li>Home</li>
		<li>Docs</li>
		<li>About</li>
	</ul>
</nav>
<slot />

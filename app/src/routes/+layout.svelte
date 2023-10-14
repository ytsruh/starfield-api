<script>
  import "../styles.css";
</script>

<header class="flex flex-wrap sm:justify-start sm:flex-nowrap z-50 w-full py-4">
  <nav
    class="max-w-[85rem] w-full mx-auto px-4 sm:flex sm:items-center sm:justify-between"
    aria-label="Global"
  >
    <div class="flex items-center justify-between">
      <a href="/">
        <img src="/starfield_logo.svg" class="h-16 w-32" alt="Starfield logo" />
      </a>
      <div class="sm:hidden">
        <button
          type="button"
          class="hs-collapse-toggle p-2 inline-flex justify-center items-center gap-2 rounded-md border font-medium bg-zinc-50 text-zinc-700 shadow-sm align-middle hover:bg-zinc-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-white transition-all text-sm"
          data-hs-collapse="#navbar-collapse-with-animation"
          aria-controls="navbar-collapse-with-animation"
          aria-label="Toggle navigation"
        >
          <svg
            class="hs-collapse-open:hidden w-5 h-5"
            width="16"
            height="16"
            fill="currentColor"
            viewBox="0 0 16 16"
          >
            <path
              fill-rule="evenodd"
              d="M2.5 12a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5zm0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5zm0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5z"
            />
          </svg>
          <svg
            class="hs-collapse-open:block hidden w-5 h-5"
            width="16"
            height="16"
            fill="currentColor"
            viewBox="0 0 16 16"
          >
            <path
              d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"
            />
          </svg>
        </button>
      </div>
    </div>
    <div
      id="navbar-collapse-with-animation"
      class="hs-collapse hidden overflow-hidden transition-all duration-300 basis-full grow sm:block"
    >
      <div class="flex flex-col gap-5 mt-5 items-center sm:flex-row sm:justify-end sm:mt-0 sm:pl-5">
        <a class="font-medium hover:text-gray-400 dark:hover:text-gray-500" href="/" aria-current="page"
          >Home</a
        >
        <a class="font-medium hover:text-gray-500" href="/docs">Docs</a>
        <a class="font-medium hover:text-gray-500" href="/login">
          <button
            type="button"
            class="py-3 px-4 inline-flex justify-center items-center gap-2 rounded-md border border-transparent font-semibold bg-zinc-200 text-zinc-950 hover:bg-transparent hover:text-zinc-50 hover:border-zinc-50 transition-all text-sm"
          >
            Login
          </button>
        </a>
      </div>
    </div>
  </nav>
</header>
<slot />

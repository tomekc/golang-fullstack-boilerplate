<script lang="ts">
    import '$lib/styles/admin-one.css';
    import favicon from '$lib/assets/favicon.svg';
    import { page } from '$app/state';

    let { children } = $props();
    const pathname = $derived(page.url.pathname);
</script>

<svelte:head>
    <link rel="icon" href={favicon}/>
</svelte:head>

<nav id="navbar-main" class="navbar is-fixed-top">
    <div class="navbar-brand">
        <a class="navbar-item is-hidden-desktop jb-aside-mobile-toggle" href={undefined} aria-label="Toggle sidebar">
            <span class="icon"><i class="mdi mdi-forwardburger mdi-24px"></i></span>
        </a>
        <div class="navbar-item has-control">
            <div class="control"><input placeholder="Search everywhere..." class="input"></div>
        </div>
    </div>
    <div class="navbar-brand is-right">
        <a class="navbar-item is-hidden-desktop jb-navbar-menu-toggle" href={undefined} data-target="navbar-menu" aria-label="Toggle menu">
            <span class="icon"><i class="mdi mdi-dots-vertical"></i></span>
        </a>
    </div>
    <div class="navbar-menu fadeIn animated faster" id="navbar-menu">
        <div class="navbar-end">
            <div class="navbar-item has-dropdown has-dropdown-with-icons has-divider has-user-avatar is-hoverable">
                <a class="navbar-link is-arrowless" href={undefined}>
                    <div class="is-user-avatar">
                        <img src="https://avatars.dicebear.com/v2/initials/john-doe.svg" alt="John Doe">
                    </div>
                    <div class="is-user-name"><span>John Doe</span></div>
                    <span class="icon"><i class="mdi mdi-chevron-down"></i></span>
                </a>
                <div class="navbar-dropdown">
                    <a class="navbar-item" href={undefined}>
                        <span class="icon"><i class="mdi mdi-account"></i></span>
                        <span>My Profile</span>
                    </a>
                    <a class="navbar-item" href={undefined}>
                        <span class="icon"><i class="mdi mdi-settings"></i></span>
                        <span>Settings</span>
                    </a>
                    <hr class="navbar-divider">
                    <a class="navbar-item" href={undefined}>
                        <span class="icon"><i class="mdi mdi-logout"></i></span>
                        <span>Log out</span>
                    </a>
                </div>
            </div>
            <a title="About" class="navbar-item has-divider is-desktop-icon-only" href="/about">
                <span class="icon"><i class="mdi mdi-help-circle-outline"></i></span>
                <span>About</span>
            </a>
            <a title="Log out" class="navbar-item is-desktop-icon-only" href={undefined}>
                <span class="icon"><i class="mdi mdi-logout"></i></span>
                <span>Log out</span>
            </a>
        </div>
    </div>
</nav>

<aside class="aside is-placed-left is-expanded">
    <div class="aside-tools">
        <div class="aside-tools-label">
            <span><b>Admin</b> One</span>
        </div>
    </div>
    <div class="menu is-menu-main">
        <p class="menu-label">General</p>
        <ul class="menu-list">
            <li>
                <a href="/" class="has-icon" class:is-active={pathname === '/'}>
                    <span class="icon"><i class="mdi mdi-desktop-mac"></i></span>
                    <span class="menu-item-label">Dashboard</span>
                </a>
            </li>
        </ul>
        <p class="menu-label">Examples</p>
        <ul class="menu-list">
            <li>
                <a href="/hello" class="has-icon" class:is-active={pathname === '/hello'}>
                    <span class="icon"><i class="mdi mdi-clock-outline"></i></span>
                    <span class="menu-item-label">Hello API</span>
                </a>
            </li>
            <li>
                <a href="/about" class="has-icon" class:is-active={pathname === '/about'}>
                    <span class="icon"><i class="mdi mdi-help-circle"></i></span>
                    <span class="menu-item-label">About</span>
                </a>
            </li>
        </ul>
    </div>
</aside>

{@render children()}

<footer class="footer">
    <div class="container-fluid">
        <div class="level">
            <div class="level-left">
                <div class="level-item">
                    © 2025, Golang-SvelteKit Boilerplate
                </div>
            </div>
        </div>
    </div>
</footer>

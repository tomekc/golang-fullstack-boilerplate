<script lang="ts">
    let { children, userInfo = 'Admin User' } = $props();
</script>

<aside id="layoutSidenav_nav" class="menu">
    <div class="sb-sidenav is-flex is-flex-direction-column">
        <div class="sb-sidenav-menu is-flex-grow-1">
            {@render children?.()}
        </div>

        <div class="sb-sidenav-footer">
            <p class="is-size-7 has-text-grey-light">Logged in as:</p>
            <p class="has-text-white">{userInfo}</p>
        </div>
    </div>
</aside>

<style>
    #layoutSidenav_nav {
        flex-basis: 225px;
        flex-shrink: 0;
        transition: transform 0.15s ease-in-out;
    }

    .sb-sidenav {
        height: 100%;
        background-color: #212529;
        color: rgba(255, 255, 255, 0.5);
    }

    .sb-sidenav-menu {
        overflow-y: auto;
    }

    .sb-sidenav-footer {
        padding: 0.75rem 1rem;
        background-color: rgba(0, 0, 0, 0.1);
    }

    /* Override Bulma .menu-list a specificity for dark sidebar */
    :global(.sb-sidenav .menu-list a) {
        color: rgba(255, 255, 255, 0.5);
        background-color: transparent;
    }

    :global(.sb-sidenav .menu-list a:hover) {
        color: #fff;
        background-color: rgba(255, 255, 255, 0.075);
    }

    :global(.sb-sidenav .menu-list a.is-active) {
        color: #fff;
        background-color: rgba(255, 255, 255, 0.1);
    }

    /* Responsive - hide sidebar on mobile */
    @media (max-width: 768px) {
        #layoutSidenav_nav {
            transform: translateX(-225px);
        }
    }
</style>

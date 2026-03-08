<script lang="ts">
    interface Props {
        appName?: string;
        version?: string;
    }

    let { appName = 'Boilerplate', version = '0.0.1' }: Props = $props();
</script>

<nav class="navbar is-dark sb-topnav">
    <div class="navbar-brand">
        <a class="navbar-item has-text-weight-bold" href="/">
            {appName}
        </a>
    </div>
    <div class="navbar-menu is-active">
        <div class="navbar-end">
            <div class="navbar-item">
                <span class="has-text-grey-light is-size-7">version {version}</span>
            </div>
        </div>
    </div>
</nav>

<style>
    .sb-topnav {
        z-index: 1039;
        height: 56px;
    }
</style>

# A Porter brig Mixin

This mixin contains the [brig](https://docs.brigade.sh/topics/brig/) CLI for use
in [Porter](https://porter.sh) bundles.

Check out Porter's [Mixin](https://porter.sh/author-bundles/#mixins)
doc on how to utilize mixins when authoring bundles.

# Install

Use the `porter` [CLI](https://porter.sh/install/) to install this
mixin via its feed URL.

```console
$ porter mixin install brig --feed-url https://vdice.github.io/porter-brig-mixin/atom.xml
installed brig mixin v0.1.0 (d605dfc)
```

Now you are ready to use the mixin in a Porter bundle!

Check out the `examples/` directory for usage examples.


# Usage

This mixin isn't yet opinionated about either v1 or v2 of the brig CLI and
doesn't have `brig`-specific action support yet.

However, one benefit of this is, as long as the `brig` version you wish to use
exists on Brigade's [releases](https://github.com/brigadecore/brigade/releases)
list, you can specify in your porter.yaml file like so:

```
mixins:
  - brig:
      clientVersion: v2.0.0-alpha.5
```

And use the mixin in an action like so:

```
install:
  - brig:
      description: Print brig version
      arguments:
        - --version
```

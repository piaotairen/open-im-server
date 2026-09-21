# ZhiLian / 勤链 fork notice

This repository is a fork of OpenIM Server for **operations and necessary
modifications only**. Business logic stays in the Java module `link-im`
of `piaotairen/link-ai-supply`.

## Provenance

| Field | Value |
|---|---|
| Upstream | https://github.com/openimsdk/open-im-server |
| Upstream default branch | `main` |
| Locked release | `v3.8.3-patch.16` (2026-03-19) |
| Upstream commit | `f6411a8a1a31d3df36f4c2b3ad28481a94141e1f` |
| Annotated tag object | `8c6bf8af683b0f739c4b7c1b3b4447cb436dcbcd` |
| This fork | https://github.com/piaotairen/open-im-server |
| Baseline tag | `zhilian-base-v3.8.3-patch.16` |
| Working branch | `zhilian` |
| Fork created | 2026-09-17 |
| Baseline locked | 2026-09-21 |
| License | Apache License 2.0 (`LICENSE`, 201 lines, AFFERO=0) |

Matching `protocol` module pin from this release `go.mod`:

`github.com/openimsdk/protocol v0.0.73-alpha.12`  
commit `0f9ddb17a47a080147ea9e635074de1b8bb6f64d`

## License obligations

Apache-2.0 allows modification, closed-source distribution, and commercial use.
We retain the original `LICENSE` (and any upstream `NOTICE` if present) and
record our changes in `NOTICE_MODIFICATIONS.md` (Apache-2.0 §4).

## Forbidden: AGPL client code

**Do not introduce any AGPL-licensed client SDK source into this fork or the
ZhiLian clients.**

Allowed reference sources (Apache-2.0 only):

| Component | Locked revision | SHA |
|---|---|---|
| `open-im-server` | `v3.8.3-patch.16` | `f6411a8a1a31d3df36f4c2b3ad28481a94141e1f` |
| `protocol` | `v0.0.73-alpha.12` | `0f9ddb17a47a080147ea9e635074de1b8bb6f64d` |
| `openim-sdk-core` | `v3.8.3-patch.3` (2025-03-07) | `93105bfe401a07a4e5c1637dd9887e637fee03a0` |
| `openim-sdk-js-wasm` | `v3.8.3-patch.3` (2025-03-07) | `d4f5ac16b8e0045e96a10641ef290419c16eba64` |

**Never open or vendor** `openim-sdk-core` / `openim-sdk-js-wasm` /
`openim-sdk-android` / `openim-sdk-flutter` / `openim-sdk-ios` at
`v3.8.3-patch.4` or later (AGPL-3.0). `openim-sdk-ios` default branch
`master` is AGPL and is **not** a reference source. iOS clients must
follow `openim-sdk-core@v3.8.3-patch.3` + `protocol` only.

When syncing upstream: merge **Apache-licensed server/protocol revisions
only**. Do not pull AGPL client trees.

Evidence copies live in `link-ai-supply/docs/vendor/openim/`.

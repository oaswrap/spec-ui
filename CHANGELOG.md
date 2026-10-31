# Changelog

All notable changes to this project are documented here.
Format based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Added

- Stoplight Elements: `TryItCredentialsPolicy` option (`omit`, `include`, `same-origin`) sets the fetch credentials
  policy of the "Try it" feature. New type `config.ElementTryItCredentialsPolicy` with constants
  `ElementTryItCredentialsPolicyOmit`, `ElementTryItCredentialsPolicyInclude` and
  `ElementTryItCredentialsPolicySameOrigin` ([#14](../../pull/14)).

### Changed

- Bump UI versions (CDN URLs, embedded assets, `Makefile`):

  | UI                 | From    | To      |
  | ------------------ | ------- | ------- |
  | Swagger UI         | 5.32.11 | 5.34.1  |
  | Stoplight Elements | 9.0.24  | 9.0.29  |
  | Scalar             | 1.63.0  | 1.76.4  |
  | ReDoc              | 2.5.3   | 3.0.1   |
  | RapiDoc            | 9.3.8   | 12.0.0  |

- **ReDoc 3**: the bundle is now an ES module. The template loads it with
  `<script type="module">` and calls `init(...)` instead of the removed global `Redoc.init(...)`.
  The `expandResponses` option is no longer emitted (removed upstream).
- **RapiDoc 12**: `config.RapiDoc.Theme` (`light`/`dark`) is now rendered as the `color-scheme` attribute.
  When unset it defaults to `light` to keep the previous look (RapiDoc 12 defaults to dark).
- **RapiDoc 12**: `BgColor` and `TextColor` are applied through the `--background` and `--foreground`
  CSS variables, because the `bg-color` and `text-color` attributes were removed upstream.

### Removed

- **Breaking:** `config.ReDoc.HideSearch`. ReDoc 3 always enables search.
- **Breaking:** `config.RapiDoc.HeaderColor`. RapiDoc 11+ has no equivalent.

### Migration

- Delete `HideSearch` from `config.ReDoc{...}` and `HeaderColor` from `config.RapiDoc{...}`; both now fail to compile.
- ReDoc 3 and RapiDoc 12 change the look and behavior of those pages (new design, deep-link format, system fonts, renamed
  theme presets). Review them visually after upgrading.

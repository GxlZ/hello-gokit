# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Added `Name` option for `Provide` to add named values to the container
  without rewriting constructors. See package documentation for more
  information.

### Changed
- `name:"..."` tags on nested Result Objects will now cause errors instead of
  being ignored.

## [1.3.0] - 2017-12-04
### Changed
- Improved messages for errors thrown by Dig under a many scenarios to be more
  informative.

## [1.2.0] - 2017-11-07
### Added
- `dig.In` and `dig.Out` now support value groups, making it possible to
  produce many values of the same type from different constructors. See package
  documentation for more information.

## [1.1.0] - 2017-09-15
### Added
- Added the `dig.RootCause` function which allows retrieving the original
  constructor error that caused an `Invoke` failure.

### Changed
- Errors from `Invoke` now attempt to hint to the user a presence of a similar
  type, for example a pointer to the requested type and vice versa.

## [1.0.0] - 2017-07-31

First stable release: no breaking changes will be made in the 1.x series.

### Changed
- `Provide` and `Invoke` will now fail if `dig.In` or `dig.Out` structs
  contain unexported fields. Previously these fields were ignored which often
  led to confusion.

## [1.0.0-rc2] - 2017-07-21
### Added
- Exported `dig.IsIn` and `dig.IsOut` so that consuming libraries can check if
  a params or return struct embeds the `dig.In` and `dig.Out` types, respectively.

### Changed
- Added variadic options to all public APIS so that new functionality can be
  introduced post v1.0.0 without introducing breaking changes.
- Functions with variadic arguments can now be passed to `dig.Provide` and
  `dig.Invoke`. Previously this caused an error, whereas now the args will be ignored.

## [1.0.0-rc1] - 2017-06-21

First release candidate.

## [0.5.0] - 2017-06-19
### Added
- `dig.In` and `dig.Out` now support named instances, i.e.:

    ```go
    type param struct {
      dig.In

      DB1 DB.Connection `name:"primary"`
      DB2 DB.Connection `name:"secondary"`
    }
    ```

### Fixed
- Structs compatible with `dig.In` and `dig.Out` may now be generated using
  `reflect.StructOf`.

## [0.4.0] - 2017-06-12
### Added
- Add `dig.In` embeddable type for advanced use-cases of specifying dependencies.
- Add `dig.Out` embeddable type for advanced use-cases of constructors
  inserting types in the container.
- Add support for optional parameters through `optional:"true"` tag on `dig.In` objects.
- Add support for value types and many built-ins (maps, slices, channels).

### Changed
- **[Breaking]** Restrict the API surface to only `Provide` and `Invoke`.
- **[Breaking]** Update `Provide` method to accept variadic arguments.

### Removed
- **[Breaking]** Remove `Must*` funcs to greatly reduce API surface area.
- Providing constructors with common returned types results in an error.

## [0.3] - 2017-05-02
### Added
- Add functionality to `Provide` to support constructor with `n` return
  objects to be resolved into the `dig.Graph`
- Add `Invoke` function to invoke provided function and insert return
  objects into the `dig.Graph`

### Changed
- Rename `RegisterAll` and `MustRegisterAll` to `ProvideAll` and
  `MustProvideAll`.

## [0.2] - 2017-03-27
### Changed
- Rename `Register` to `Provide` for clarity and to recude clash with other
  Register functions.
- Rename `dig.Graph` to `dig.Container`.

### Removed
- Remove the package-level functions and the `DefaultGraph`.

## 0.1 - 2017-03-23

Initial release.

[Unreleased]: https://github.com/uber-go/dig/compare/v1.3.0...HEAD
[1.3.0]: https://github.com/uber-go/dig/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/uber-go/dig/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/uber-go/dig/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/uber-go/dig/compare/v1.0.0-rc2...v1.0.0
[1.0.0-rc2]: https://github.com/uber-go/dig/compare/v1.0.0-rc1...v1.0.0-rc2
[1.0.0-rc1]: https://github.com/uber-go/dig/compare/v0.5.0...v1.0.0-rc1
[0.5.0]: https://github.com/uber-go/dig/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/uber-go/dig/compare/v0.3...v0.4.0
[0.3]: https://github.com/uber-go/dig/compare/v0.2...v0.3
[0.2]: https://github.com/uber-go/dig/compare/v0.1...v0.2
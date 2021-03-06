// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Base: Functions and other definitions useful with multiple packages.
//
// The book Astrononomical Algorithsms begins with an unnumbered chapter
// titled "Some Symbols and Abbreviations."  In addition to a list of symbols
// and abbreviations are a few paragraphs introducing sexagesimal notation.
// Chapter 1, Hints and Tips contains additional information about sexagesimal
// numbers.  It made sense to combine these in one package.  Also here
// are various definitions and support functions useful in multiple chapters.
//
// Decimal Symbols
//
// Described on p.6 is a convention for placing a
// unit symbol directly above the decimal point of a decimal number.
// This can be done with Unicode by replacing the decimal point with
// the unit symbol and "combining dot below," u+0323.  The function
// DecSymCombine here performs this substitution.  Of course this only
// works to the extent that software can render the combining character.
// For cases where rendering software fails badly, DecSymAdd is provided
// as a compromise.  It does not use the combining dot but simply places
// the unit symbol ahead of the decimal point.  Numbers modified with either
// function can be returned to their original form with DecSymStrip.
//
// Sexagesimal types
//
// Not described in AA, but of great use are four types commonly expressed
// in sexagesimal format: Angle, HourAngle, RA, and Time.
// The underlying type of each is float64.  The unit for Angle, HourAngle,
// and RA is radians.  The unit for Time is seconds.
// Each type has a constructor that takes sexegesimal components.  Each type
// also has a method, Rad or Seconds, that simply returns the underlying type.
// Being based on float64s, these types are relatively efficient.
//
// Custom formatters
//
// Parallel to the four types just described are four types with custom
// formatters that produce sexagesimal formattng.  These types are FmtAngle,
// FmtHourAngle, FmtRA, and FmtTime.  These types are structs with methods
// that have pointer receivers.  There is more overhead with these types
// than with the basic Angle, HourAngle, RA, and Time types.
//
// The syntax of a format specifier is
//  %[flags][width][.precision]verb
//
// Verbs are s, d, c, x, and v.  The meanings are different than for
// common Go types.  Given an Angle equivalent to 1.23 seconds,
//  %.2s formats as 1.23″   (s for standard formatting)
//  %.2d formats as 1″.23   (d for decimal symbol, as in DecSymAdd)
//  %.2c formats as 1″̣23    (c for combining dot, as in DecSymCombine)
//  %.2x formats as 123     (x for space, suppresses unit symbols and decimal point)
//  %v formats the same as %s
//
// The following flags are supported:
//  + always print leading sign
//  ' ' (space) leave space for elided sign
//  # display all three segments, even if 0
//  0 pad all segments with leading zeros
//
// A + flag takes precedence over a ' ' (space) flag.
// The # flag forces all formatted strings to have three numeric components,
// an hour or degree, a minute, and a second.  Without the # flag, small vaues
// will have zero values of hours, degrees, or minutes elided.
// The 0 flag pads with leading zeros on minutes and seconds, and if a
// width is specfied, leading zeros on the first segment as well.
// For the RA type, sign formatting flags '+' and ' ' are ignored.
//
// Width specifies the number of digits in the most significant segment,
// degrees or hours (not the total width of all three segments.)
//
// Precision specifies the number of places past the decimal point
// of the last (seconds) segment.  There are two magic numbers for precision
// however:  62 means to round to the nearest minute and not show seconds
// at all.  64 means to round to the nearest degree or hour and not show
// minutes or seconds at all.
//
// To ensure fixed width output, use one of the + or ' ' (space) flags,
// use the 0 flag, and use a width.
//
// The symbols used for degrees, minutes, and seconds for the FmtAngle type
// are taken from the package variable DMSRunes.  The symbols for
// hours, minutes, and seconds for the FmtHourAngle, FmtRA, and FmtTime
// types are taken from HMSRunes.
//
// Width Errors
//
// For various types of overflow, the custom formatters emit all asterisks
// "*************" and leave an exact error in the WidthError field of the
// type.
//
// Precision is limited to the range [0,15].  Values outside of that range
// will cause this overflow condition.
//
// Precision of 15 is possible only for angles less than a few arc seconds.
// As angle values increase, fewer digits of precision are possible.  At one
// degree, you can get 12 digits of precision, at 360 degrees, you can get 9.
// An angle too large for the specified precision causes overflow.
//
// If you specifiy width, the first segment, degrees or hours, must fit in the
// specified width.  Larger values cause overflow.
//
// +Inf, -Inf, and NaN always cause overflow.
//
// Bessellian and Julian Year
//
// Chapter 21, Precession actually contains these defintions.  They are moved
// here because of their general utility.
//
// Chapter 22, Nutation contains the function for Julian centuries since J2000.
//
// Phase angle functions
//
// Two functions, Illuminated and Limb, concern the illumnated phase of a body
// and are given in two chapters, 41 an 48.  They are collected here because
// the identical functions apply in both chapters.
//
// General purpose math functions
//
// SmallAngle is recommended in chapter 17, p. 109.
//
// PMod addresses the issue on p. 7, chapter 1, in the section "Trigonometric
// functions of large angles", but the function is not written to be specific
// to angles and so has more general utility.
//
// Horner is described on p. 10, chapter 1.
//
// FloorDiv and FloorDiv64 are optimizations for the INT function described
// on p. 60, chapter 7.
package base

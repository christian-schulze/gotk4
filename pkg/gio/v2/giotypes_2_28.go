// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// CancellableSourceFunc: this is the function type of the callback used for the
// #GSource returned by g_cancellable_source_new().
type CancellableSourceFunc func(cancellable context.Context) (ok bool)

// PollableSourceFunc: this is the function type of the callback used for the
// #GSource returned by g_pollable_input_stream_create_source() and
// g_pollable_output_stream_create_source().
type PollableSourceFunc func(pollableStream *coreglib.Object) (ok bool)
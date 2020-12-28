/*
 *    Copyright (C) 2014-2017 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

package webbee

import (
	"github.com/muesli/beehive/bees"
)

// WebBeeFactory is a factory for WebBees.
type WebBeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *WebBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := WebBee{
		Bee: bees.NewBee(name, factory.ID(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// ID returns the ID of this Bee.
func (factory *WebBeeFactory) ID() string {
	return "webbee"
}

// Name returns the name of this Bee.
func (factory *WebBeeFactory) Name() string {
	return "HTTP Server"
}

// Description returns the description of this Bee.
func (factory *WebBeeFactory) Description() string {
	return "Reacts to HTTP requests"
}

// Image returns the filename of an image for this Bee.
func (factory *WebBeeFactory) Image() string {
	return factory.ID() + ".png"
}

// LogoColor returns the preferred logo background color (used by the admin interface).
func (factory *WebBeeFactory) LogoColor() string {
	return "#223f5e"
}

// Options returns the options available to configure this Bee.
func (factory *WebBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		{
			Name:        "address",
			Description: "Which addr to listen on, eg: 0.0.0.0:12345",
			Type:        "address",
			Mandatory:   true,
		},
	}
	return opts
}

// Events describes the available events provided by this Bee.
func (factory *WebBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "get",
			Description: "A GET call was received by the HTTP server",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "remote_addr",
					Description: "IP of the caller",
					Type:        "address",
				},
				{
					Name:        "url",
					Description: "Request URL",
					Type:        "url",
				},
				{
					Name:        "query_params",
					Description: "Map of query parameters received from caller",
					Type:        "map",
				},
				{
					Name:        "data",
					Description: "Raw request data",
					Type:        "string",
				},
				{
					Name:        "json",
					Description: "JSON map received from caller",
					Type:        "map",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "post",
			Description: "A POST call was received by the HTTP server",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "remote_addr",
					Description: "IP of the caller",
					Type:        "address",
				},
				{
					Name:        "url",
					Description: "Request URL",
					Type:        "url",
				},
				{
					Name:        "query_params",
					Description: "Map of query parameters received from caller",
					Type:        "map",
				},
				{
					Name:        "data",
					Description: "Raw request data",
					Type:        "string",
				},
				{
					Name:        "json",
					Description: "JSON map received from caller",
					Type:        "map",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "put",
			Description: "A PUT call was received by the HTTP server",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "remote_addr",
					Description: "IP of the caller",
					Type:        "address",
				},
				{
					Name:        "url",
					Description: "Request URL",
					Type:        "url",
				},
				{
					Name:        "query_params",
					Description: "Map of query parameters received from caller",
					Type:        "map",
				},
				{
					Name:        "data",
					Description: "Raw request data",
					Type:        "string",
				},
				{
					Name:        "json",
					Description: "JSON map received from caller",
					Type:        "map",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "patch",
			Description: "A PATCH call was received by the HTTP server",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "remote_addr",
					Description: "IP of the caller",
					Type:        "address",
				},
				{
					Name:        "url",
					Description: "Request URL",
					Type:        "url",
				},
				{
					Name:        "query_params",
					Description: "Map of query parameters received from caller",
					Type:        "map",
				},
				{
					Name:        "data",
					Description: "Raw request data",
					Type:        "string",
				},
				{
					Name:        "json",
					Description: "JSON map received from caller",
					Type:        "map",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "delete",
			Description: "A DELETE call was received by the HTTP server",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "remote_addr",
					Description: "IP of the caller",
					Type:        "address",
				},
				{
					Name:        "url",
					Description: "Request URL",
					Type:        "url",
				},
				{
					Name:        "query_params",
					Description: "Map of query parameters received from caller",
					Type:        "map",
				},
				{
					Name:        "data",
					Description: "Raw request data",
					Type:        "string",
				},
				{
					Name:        "json",
					Description: "JSON map received from caller",
					Type:        "map",
				},
			},
		},
	}
	return events
}

func init() {
	f := WebBeeFactory{}
	bees.RegisterFactory(&f)
}

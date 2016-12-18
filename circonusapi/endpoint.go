// Copyright 2016 Alem Abreha <alem.abreha@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// API endpoint setting
package circonusapi

import "os"

// Circonus API URL from env
var CirconusURL = os.Getenv("CIRCONUS_API_URL")

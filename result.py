#!/usr/bin/env python
#
# Copyright 2007 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
import cgi
import datetime
import random
import wsgiref.handlers

from google.appengine.ext import db
from google.appengine.api import users
from google.appengine.ext import webapp

seq = ("""
           <head>
             <title>Yes Coffee is Brewing!!</title>
           </head>
           <body bgcolor=green>
           %s
           <center>%s</center>
           </br>
           <center>
           <img src="/img/coffee-works.jpg">
           <h1>YES! The Coffee is ON!</h1>
           </center>
           <center>%s</center>
           </body>
           </html>""",
      """
            <head>
              <title>NO Coffee is BROKEN!!</title>
            </head>
            <body bgcolor=red>
            %s
            <center>%s</center>
            </br>
            <center>
            <img src="/img/no-coffee.jpg">
            <h1>NO! The Coffee is BROKEN!</h1>
            </center>
            <center>%s</center>
            </body>
            </html>""")

class MainPage(webapp.RequestHandler):
  def get(self):
    
    spacer = "<br>\n" * 10
    top_ad = """<script type="text/javascript"><!--
    google_ad_client = "ca-pub-4109530843345987";
    /* top-coffee-machine */
    google_ad_slot = "2520183523";
    google_ad_width = 728;
    google_ad_height = 90;
    //-->
    </script>
    <script type="text/javascript"
    src="http://pagead2.googlesyndication.com/pagead/show_ads.js">
    </script>
    """
    bot_ad = """<script type="text/javascript"><!--
    google_ad_client = "ca-pub-4109530843345987";
    /* bottom-coffee */
    google_ad_slot = "2064675845";
    google_ad_width = 728;
    google_ad_height = 90;
    //-->
    </script>
    <script type="text/javascript"
    src="http://pagead2.googlesyndication.com/pagead/show_ads.js">
    </script>"""

    self.response.out.write('<html>')


    self.response.out.write(random.choice(seq) % (spacer, top_ad, bot_ad))


application = webapp.WSGIApplication([
  ('/', MainPage),
], debug=True)


def main():
  wsgiref.handlers.CGIHandler().run(application)


if __name__ == '__main__':
  main()

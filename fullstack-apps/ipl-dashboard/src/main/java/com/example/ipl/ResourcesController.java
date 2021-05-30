package com.example.ipl;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
public class ResourcesController {
  /*
   * Redirects routes to FrontEnd
   */
  @RequestMapping(value = {"/teams/**"})
  public String redirectApi() {
        return "forward:/";
    }
}

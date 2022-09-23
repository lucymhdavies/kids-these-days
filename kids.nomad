job "kids-these-days" {
  datacenters = ["davnet"]
  type        = "batch"
  namespace   = "twitter"

  periodic {
    cron             = "@hourly"
    prohibit_overlap = true
  }

  group "bot" {
    count = 1

    task "run" {
      driver = "docker"

      artifact {
        source      = "git::https://github.com/lucymhdavies/kids-these-days"
        destination = "local/repo"
      }

      vault {
        policies = [
          "twitter-kids-these-days",
        ]
      }


      template {
        destination = "secrets/env"

        data = <<-EOH
        {{ with secret "kv/data/twitter/bots/kids-these-days" }}
        ACCESS_SECRET={{ .Data.data.ACCESS_SECRET }}
        ACCESS_TOKEN={{ .Data.data.ACCESS_TOKEN }}
        CONSUMER_KEY={{ .Data.data.CONSUMER_KEY }}
        CONSUMER_SECRET={{ .Data.data.CONSUMER_SECRET }}
        {{ end }}
        EOH

        env = true
      }

      config {
        image = "golang:1.19"

        volumes = [
          "local/repo:/tmp/workdir"
        ]
        work_dir = "/tmp/workdir"

        args = [
          "go", "run", "."
        ]
      }

      resources {
        # While it only needs to be online for a few seconds, it does seem to use quite a bit of CPU
        cpu = 3 * 1024

        # Memory at least is minimal
        memory = 128
      }

    }
  }
}


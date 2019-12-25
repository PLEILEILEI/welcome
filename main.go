package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"welcome/probe"

	"github.com/gin-gonic/gin"
)

func main() {
	cmd := &cobra.Command{
		Use:"welcome",
	}
	cmd.AddCommand(newLiveCommand(),newRunCommand())
	cmd.Execute()

}
func newRunCommand() *cobra.Command {
	return  &cobra.Command{
		Use:"run",
		Run: func(cmd *cobra.Command, args []string) {
			if err := probe.Create(); err != nil {
				panic(fmt.Sprintf("liveness probe init failed | err:%+v\n",err))
			}
			router := gin.Default()
			router.Use(func(c *gin.Context) {
				if c.Request.URL.Path == "/ready" {
					c.AbortWithStatus(http.StatusOK)
				}
			})

			router.GET("/welcome", func(c *gin.Context) {
				c.String(http.StatusOK, "Welcome Gin Server, Now is %s", time.Now().String())
			})

			srv := &http.Server{
				Addr:    ":8888",
				Handler: router,
			}

			go func() {
				// service connectionscd
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()

			quit := make(chan os.Signal)
			// kill (no param) default send syscall.SIGTERM
			// kill -2 is syscall.SIGINT
			// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			if err := probe.Remove(); err != nil {
				log.Fatal("liveness probe remove failed | err:",err)
			}
			// catching ctx.Done(). timeout of 5 seconds.
			select {
			case <-ctx.Done():
				log.Println("timeout of 5 seconds.")
			}
			log.Println("Server exiting")
		},
	}
}
func newLiveCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:  "live",
		Run: func(cmd *cobra.Command, args []string) {
			if probe.Exists() {
				println("liveness ok")
				os.Exit(0)
			}
			println("liveness failed")
			os.Exit(1)
		},
	}
	return cmds
}

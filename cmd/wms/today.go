package wms

import (
	"context"
	"fmt"
	"log"

	"github.com/MESMUR/wms/pkg/events"
	"github.com/MESMUR/wms/pkg/initialize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var defToday bool

var todayCmd = &cobra.Command{
	Use:     "today",
	Aliases: []string{"tod"},
	Short:   "Gets today's schedule!",
	Long:    "Gets today's schedule for the given calendar!",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var calendarID string
		if len(args) > 0 {
			calendarID = args[0]
		} else {
			calendarID = fmt.Sprint(viper.Get("calendar_name"))
		}

		if defToday {
			viper.Set("calendar_name", calendarID)
			viper.WriteConfig()
		}

		ctx := context.Background()

		client := initialize.GetClient()
		srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))

		if err != nil {
			log.Fatalf("Unable to retrieve Calendar client: %v", err)
		}

		events.GetEvents(srv, calendarID, 0)
	},
}

func init() {
	todayCmd.Flags().
		BoolVarP(&defToday, "default", "d", false, "Sets the provided calendar as the default")
	rootCmd.AddCommand(todayCmd)
}

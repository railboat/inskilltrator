// Package intent exports the names of the standard intents available to Amazon Skills.
package intent

// These are the standard Amazon Intents retrieved from
// https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents
const (
	HELP        = "AMAZON.HelpIntent"
	CANCEL      = "AMAZON.CancelIntent"
	LOOP_OFF    = "AMAZON.LoopOffIntent"
	LOOP_ON     = "AMAZON.LoopOnIntent"
	NEXT        = "AMAZON.NextIntent"
	NO          = "AMAZON.NoIntent"
	PAUSE       = "AMAZON.PauseIntent"
	PREVIOUS    = "AMAZON.PreviousIntent"
	REPEAT      = "AMAZON.RepeatIntent"
	RESUME      = "AMAZON.ResumeIntent"
	SHUFFLE_OFF = "AMAZON.ShuffleOffIntent"
	SHUFFLE_ON  = "AMAZON.ShuffleOnIntent"
	RESTART     = "AMAZON.StartOverIntent"
	STOP        = "AMAZON.StopIntent"
	YES         = "AMAZON.YesIntent"
)

package raymond

import "testing"

// A block helper called with a param that resolves to nothing used to receive a
// typed-nil *Options and dereference it.
func TestBlockHelperWithMissingParam(t *testing.T) {
	tests := map[string]string{
		"if":        `{{#if user.first_name user.last_name}}hi{{/if}}`,
		"unless":    `{{#unless user.first_name user.last_name}}hi{{/unless}}`,
		"with":      `{{#with user.first_name user.last_name}}hi{{/with}}`,
		"inline if": `{{if user.first_name user.last_name}}`,
	}

	ctx := map[string]any{"user": map[string]any{"first_name": "Ana"}}

	for name, source := range tests {
		t.Run(name, func(t *testing.T) {
			tpl, err := Parse(source)
			if err != nil {
				t.Fatalf("Parse(%q) = %v, want no error", source, err)
			}

			got, err := tpl.Exec(ctx)
			if err == nil {
				t.Fatalf("Exec() = %q, nil, want an error", got)
			}
		})
	}
}

package sdk

import "context"

var _ = (*pipes)(nil)

type pipes struct {
	client *Client
}

func (v *pipes) Create(ctx context.Context, id AccountObjectIdentifier, opts *PipeCreateOptions) error {
	if opts == nil {
		opts = &PipeCreateOptions{}
	}
	opts.name = id
	return validateAndExec(v.client, ctx, opts)
}

func (v *pipes) Alter(ctx context.Context, id AccountObjectIdentifier, opts *PipeAlterOptions) error {
	if opts == nil {
		opts = &PipeAlterOptions{}
	}
	opts.name = id
	return validateAndExec(v.client, ctx, opts)
}

func (v *pipes) Drop(ctx context.Context, id AccountObjectIdentifier, opts *PipeDropOptions) error {
	if opts == nil {
		opts = &PipeDropOptions{}
	}
	opts.name = id
	return validateAndExec(v.client, ctx, opts)
}

func (v *pipes) Show(ctx context.Context, opts *PipeShowOptions) ([]*Pipe, error) {
	if opts == nil {
		opts = &PipeShowOptions{}
	}
	dbRows, err := validateAndQuery[pipeDBRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}

	resultList := make([]*Pipe, len(*dbRows))
	for i, row := range *dbRows {
		resultList[i] = row.toPipe()
	}

	return resultList, nil
}

func (v *pipes) ShowByID(ctx context.Context, id AccountObjectIdentifier) (*Pipe, error) {
	pipes, err := v.Show(ctx, &PipeShowOptions{
		Like: &Like{
			Pattern: String(id.Name()),
		},
	})
	if err != nil {
		return nil, err
	}

	for _, pipe := range pipes {
		if pipe.ID().name == id.Name() {
			return pipe, nil
		}
	}
	return nil, ErrObjectNotExistOrAuthorized
}

func (v *pipes) Describe(ctx context.Context, id SchemaObjectIdentifier) (*PipeDetails, error) {
	//TODO implement me
	panic("implement me")
}

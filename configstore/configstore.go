package configstore

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/consul/api"
	"github.com/milossimic/rest/tracer"
)

type ConfigStore struct {
	cli *api.Client
}

func (ps *ConfigStore) AddIdempotencyKey(ctx context.Context, id string) {

	span:=tracer.StartSpanFromContext(ctx, "AddIdempotencyKey")
	defer span.Finish()
	kv := ps.cli.KV()

	sid := createNewIdempotencyKey(id)

	data, _ := json.Marshal(true)

	c := &api.KVPair{Key: sid, Value: data}
	_, _ = kv.Put(c, nil)
}

func (ps *ConfigStore) IdempotencyKeyExists(ctx context.Context, id string) (bool, error) {

	span:= tracer.StartSpanFromContext(ctx, "IdempotencyKeyExists")
	defer span.Finish()

	kv := ps.cli.KV()
	data, _, err := kv.Get(createNewIdempotencyKey(id), nil)
	if err != nil || data == nil {
		tracer.LogError(span, err)
		return false, err
	}
	return true, nil

	//if err != nil || data == nil {
	//	return false
	//}
	//return true

	//pair, _, err := kv.Get(constructKey(id), nil)
	//if err != nil {
	//	tracer.LogError(span, err)
	//	return nil, err
	//}
}

func New() (*ConfigStore, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConfigStore{
		cli: client,
	}, nil
}

func (ps *ConfigStore) Get(ctx context.Context, id string, version string) ([]*Config, error) {

	span := tracer.StartSpanFromContext(ctx, "Get")
	defer span.Finish()
	kv := ps.cli.KV()

	data, _, err := kv.List(constructKey(id, version, ""), nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	configs := []*Config{}
	for _, pair := range data {
		post := &Config{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}
		configs = append(configs, post)
	}
	return configs, nil
}

func (ps *ConfigStore) GetAll(ctx context.Context) ([]*Config, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()
	kv := ps.cli.KV()

	data, _, err := kv.List(all, nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	posts := []*Config{}
	for _, pair := range data {
		post := &Config{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (ps *ConfigStore) Post(ctx context.Context, post *Config) (*Config, error) {
	span := tracer.StartSpanFromContext(ctx, "Post")
	defer span.Finish()
	kv := ps.cli.KV()

	sid, rid := generateKey(post.Version, "")
	post.Id = rid

	data, err := json.Marshal(post)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return post, nil
	// return nil, nil
}

func (ps *ConfigStore) Delete(ctx context.Context, id string, version string) (map[string]string, error) {
	span := tracer.StartSpanFromContext(ctx, "Delete")
	defer span.Finish()

	kv := ps.cli.KV()
	_, err := kv.DeleteTree(constructKey(id, version, ""), nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return map[string]string{"Deleted": id}, nil
}

func (ps *ConfigStore) GetPostsByLabels(ctx context.Context, id string, version string, labels string) ([]*Config, error) {
	span := tracer.StartSpanFromContext(ctx, "GetPostByLabels")
	defer span.Finish()

	kv := ps.cli.KV()

	data, _, err := kv.List(constructKey(id, version, labels), nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	posts := []*Config{}

	for _, pair := range data {
		post := &Config{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return posts, nil
}

func (ps *ConfigStore) AddConfigurationGroup(ctx context.Context, configs []*Config) ([]*Config, error) {
	span := tracer.StartSpanFromContext(ctx, "AddCofigurationGroup")
	defer span.Finish()

	kv := ps.cli.KV()

	for i := 0; i < len(configs); i++ {
		sid, rid := generateKey(configs[i].Version, "")
		configs[i].Id = rid

		data, err := json.Marshal(configs[i])
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}

		p := &api.KVPair{Key: sid, Value: data}
		_, err = kv.Put(p, nil)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}
	}

	return configs, nil
	// return nil, nil
}

func (ps *ConfigStore) GetGoupById(ctx context.Context, group_id string, group_version string) ([]*Config, error) {
	span := tracer.StartSpanFromContext(ctx, "GetGroupById")
	defer span.Finish()

	kv := ps.cli.KV()

	data, _, err := kv.List(all, nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	configs := []*Config{}
	for _, pair := range data {
		post := &Config{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}
		if post.Group_Id == group_id {
			if post.Group_Version == group_version {
				configs = append(configs, post)
			}
		}
	}

	return configs, nil
}

func (ps *ConfigStore) DeleteGoupById(ctx context.Context, group_id string, group_version string) ([]*Config, error) {

	span := tracer.StartSpanFromContext(ctx, "DeleteGroupById")
	defer span.Finish()

	kv := ps.cli.KV()

	data, _, err := kv.List(all, nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	configs := []*Config{}
	for _, pair := range data {
		post := &Config{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}
		if post.Group_Id == group_id {
			if post.Group_Version == group_version {
				configs = append(configs, post)
				_, err := kv.DeleteTree(constructKey(post.Id, post.Version, ""), nil)
				if err != nil {
					tracer.LogError(span, err)
					return nil, err
				}
			}
		}
	}

	return configs, nil
}

func (ps *ConfigStore) GetAllGroups(ctx context.Context) ([]*Config, error) {

	span := tracer.StartSpanFromContext(ctx, "GetAllGroups")
	defer span.Finish()

	kv := ps.cli.KV()

	data, _, err := kv.List(all, nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	configs := []*Config{}
	for _, pair := range data {
		config := &Config{}
		err = json.Unmarshal(pair.Value, config)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}
		if config.Group_Id != "" {
			if config.Group_Version != "" {
				configs = append(configs, config)

			}
		}
	}

	return configs, nil
}

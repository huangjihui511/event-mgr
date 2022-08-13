COMMIT="latest"

db:
	docker build -t jihui-event-mgr:$(COMMIT) .
dr:
	docker run  -d -p 8080:8080  jihui-event-mgr:$(COMMIT)
acr:
	docker login jihui.azurecr.io
	docker tag jihui-event-mgr jihui.azurecr.io/jihui-event-mgr
	docker push jihui.azurecr.io/jihui-event-mgr
.PHONY: clean
clean:

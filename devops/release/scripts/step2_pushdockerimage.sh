ACR_NAME='asorelease'

echo "ACR_NAME: $ACR_NAME"
echo "AZURESERVICEOPERATOR_IMAGE: $AZURESERVICEOPERATOR_IMAGE"
echo "AZURESERVICEOPERATOR_IMAGE_BASE: $AZURESERVICEOPERATOR_IMAGE_BASE"
echo "AZURESERVICEOPERATOR_IMAGE_LATEST: $AZURESERVICEOPERATOR_IMAGE_LATEST"
echo "AZURESERVICEOPERATOR_IMAGE_PUBLIC: $AZURESERVICEOPERATOR_IMAGE_PUBLIC"
echo "AZURESERVICEOPERATOR_IMAGE_VERSION: $AZURESERVICEOPERATOR_IMAGE_VERSION"


az acr login --name $ACR_NAME
docker pull $ACR_NAME.azurecr.io/$AZURESERVICEOPERATOR_IMAGE
docker tag $ACR_NAME.azurecr.io/$AZURESERVICEOPERATOR_IMAGE $ACR_NAME.azurecr.io/$AZURESERVICEOPERATOR_IMAGE_LATEST
docker tag $ACR_NAME.azurecr.io/$AZURESERVICEOPERATOR_IMAGE $ACR_NAME.azurecr.io/$AZURESERVICEOPERATOR_IMAGE_PUBLIC
docker image ls
docker push $ACR_NAME.azurecr.io/$AZURESERVICEOPERATOR_IMAGE_PUBLIC
docker push $ACR_NAME.azurecr.io/$AZURESERVICEOPERATOR_IMAGE_LATEST

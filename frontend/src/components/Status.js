import React from 'react';
import TopNavbar from './TopNavbar.js';
import axios from 'axios';
import '../style/CarsCatalogStyle.css';
import Urls from '../util/Urls.js';
import {
  Button,
  ButtonToolbar,
  FormGroup,
  ControlLabel,
  FormControl,
  Popover,
  Tooltip,
  Modal
} from 'react-bootstrap/lib';

export default class Status extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      show: false,
      showResa: false,
      cars: [],
      agencies: [],
      filtersModel: '',
      filtersPowerMin: '',
      filtersPowerMax: '',
      filtersLength: '',
      filtersWidth: '',
      filtersHeight: '',
      filtersPoids: '',
      filtersImmat: '',
      filtersYearA: '',
      filtersYearB: '',
      filtersPicture: ''
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleShow = this.handleShow.bind(this);
    this.handleClose = this.handleClose.bind(this);
    this.handleShowResa = this.handleShowResa.bind(this);
    this.handleCloseResa = this.handleCloseResa.bind(this);
    this.addCar = this.addCar.bind(this);
    // this.deleteCar = this.deleteCar.bind(this);
    this.handleChangeModel = this.handleChangeModel.bind(this);
    this.handleChangeYearA = this.handleChangeYearA.bind(this);
    this.handleChangeYearB = this.handleChangeYearB.bind(this);
    this.handleChangePowerMin = this.handleChangePowerMin.bind(this);
    this.handleChangePowerMax = this.handleChangePowerMax.bind(this);
    this.handleChangeWidth = this.handleChangeWidth.bind(this);
    this.handleChangeLength = this.handleChangeLength.bind(this);
    this.handleChangeHeight = this.handleChangeHeight.bind(this);
    this.handleChangePoids = this.handleChangePoids.bind(this);
    this.handleChangeImmat = this.handleChangeImmat.bind(this);
    this.handleChangePicture = this.handleChangePicture.bind(this);
    this.resetFilters = this.resetFilters.bind(this);
    this.search = this.search.bind(this);
  }

  componentDidMount() {
    // Initialisation de page : Liste des véhicules
    axios.get(`${Urls.api}/status`).then(response => {
      this.setState({ cars: response.data });
    });
    // Liste des agences
    axios.get('https://reqres.in/api/users?page=1').then(response => {
      this.setState({ agencies: response.data.data });
    });
  }

  search() {
    axios
      .get(
        `${Urls.api}/car?minYear=` +
          this.state.filtersYearA +
          '&maxYear=' +
          this.state.filtersYearB +
          '&model=' +
          this.state.filtersModel +
          '&minPower=' +
          this.state.filtersPowerMin +
          '&maxPower=' +
          this.state.filtersPowerMax +
          '&length=' +
          this.state.filtersLength +
          '&width=' +
          this.state.filtersWidth +
          '&height=' +
          this.state.filtersHeight
      )
      .then(response => {
        this.setState({ cars: response.data });
      });
    this.resetFilters();
  }

  handleChange(event) {
    this.setState({ filtersPowerMin: event.target.value });
  }

  // Ferme la Pop up d'ajout de véhicule
  handleClose() {
    this.resetFilters();
    this.setState({ show: false });
  }

  // ouvre la Pop up d'ajout de véhicule
  handleShow() {
    this.resetFilters();
    this.setState({ show: true });
  }

  // Ferme la Pop up de resa
  handleCloseResa() {
    this.setState({ showResa: false });
  }

  // ouvre la Pop up de resa
  handleShowResa() {
    this.setState({ showResa: true });
  }

  resetFilters() {
    this.setState({ filtersModel: '' });
    this.setState({ filtersYearA: '' });
    this.setState({ filtersYearB: '' });
    this.setState({ filtersPowerMin: '' });
    this.setState({ filtersPowerMax: '' });
    this.setState({ filtersLength: '' });
    this.setState({ filtersWidth: '' });
    this.setState({ filtersHeight: '' });
    this.setState({ filtersPoids: '' });
    this.setState({ filtersImmat: '' });
    this.setState({ filtersHeight: '' });
  }

  // Insere dans la Bdd le nouveau véhicule
  addCar() {
    let newCar = [];
    newCar['Immatriculation'] = this.state.filtersImmat;
    newCar['ModelName'] = this.state.filtersModel;
    newCar['Year'] = this.state.filtersYearA;
    newCar['Weight'] = this.state.filtersPoids;
    newCar['Power'] = this.state.filtersPowerMin;
    newCar['Length'] = this.state.filtersLength;
    newCar['Width'] = this.state.filtersWidth;
    newCar['Height'] = this.state.filtersHeight;
    newCar['Picture'] =
      'https://images.caradisiac.com/images/0/6/0/5/180605/S0-plus-belle-voiture-de-l-annee-troisieme-elimination-voici-les-finalistes-615440.jpg';
    newCar['AgencyID'] = 1;
    /**
    let newCar2 = {
      ModelName: 'eeeee',
      Power: 7,
      Length: 7,
      Width: 7,
      Height: 7,
      filtersPoids: 7,
      Immatriculation: '4gjoi',
      filtersYearA: 2014,
      agenceID: 1,
      Year: 2014,
      Weight: 7,
      Picture: 'url',
      AgencyID: 1
    }; */
    let newCar2 = {
      Immatriculation: this.state.filtersImmat,
      ModelName: this.state.filtersModel,
      Year: parseInt(this.state.filtersYearA),
      Weight: 42,
      Power: parseInt(this.state.filtersPowerMin),
      Length: parseInt(this.state.filtersLength),
      Width: parseInt(this.state.filtersWidth),
      Height: parseInt(this.state.filtersHeight),
      Picture: this.state.filtersPicture,
      AgencyID: 1
    };

    let a = JSON.stringify(newCar2);
    axios.post(`${Urls.api}/car`, newCar2).then(response => {
      axios.get(`${Urls.api}/car`).then(response => {
        this.setState({ cars: response.data });
        this.handleClose();
      });
    });
  }

  // Supprime dans la Bdd le véhicule
  deleteCar(id) {
    // axios.delete(`${Urls.api}/car/${id}`).then(response => {
    // this.setState({ agencies: response.data.data });
    //});
    console.log(id);
  }

  /**
   * Liste des véhicules
   */
  tabCars() {
    let that = this;
    return this.state.cars.map(function(car, index) {
      return (
        <div className="car">
          <p>{car.Libelle}</p>

          <Button
            bsStyle="danger"
            className="carDelete"
            onClick={that.deleteCar(car.ID)}
          >
            Supprimer
          </Button>
        </div>
      );
    });
  }

  /**
   * Filtres - agences
   */
  /** agenciesfilter() {
    return this.state.cars.map(function(car, index) {
      return (
        <div>
          <input type="checkbox" value="{car.name}" id="searchColumn" />
          <label for="scales">{car.first_name}</label>
        </div>
      );
    });
  } */

  /**
   * Filtres - status
   */
  /** statusfilter() {
    return this.state.cars.map(function(car, index) {
      return (
        <div>
          <input type="checkbox" value="{car.name}" id="searchColumn" />
          <label for="scales">{car.first_name}</label>
        </div>
      );
    });
  } */

  /**
   * Page véhicules
   */
  render() {
    return (
      <div>
        <TopNavbar />
        <div
          className="static-modal"
          style={{ display: this.state.show ? 'inline' : 'none' }}
        >
          <Modal.Dialog>
            <Modal.Header>
              <Modal.Title>Ajouter un statut</Modal.Title>
            </Modal.Header>
            <input
              type="text"
              placeholder="nom modèle"
              value={this.state.filtersModel}
              onChange={this.handleChangeModel}
            />
            <input
              type="text"
              placeholder="Année"
              value={this.state.filtersYearA}
              onChange={this.handleChangeYearA}
            />
            <input
              type="text"
              placeholder="Poids"
              value={this.state.filtersPoids}
              onChange={this.handleChangePoids}
            />
            <input
              type="text"
              placeholder="Puissance"
              value={this.state.filtersPowerMin}
              onChange={this.handleChangePowerMin}
            />
            <input
              type="text"
              placeholder="Longueur"
              value={this.state.filtersLength}
              onChange={this.handleChangeLength}
            />
            <input
              type="text"
              placeholder="Largeur"
              value={this.state.filtersWidth}
              onChange={this.handleChangeWidth}
            />
            <input
              type="text"
              placeholder="Hauteur"
              value={this.state.filtersHeight}
              onChange={this.handleChangeHeight}
            />
            <input
              type="text"
              placeholder="Image"
              value={this.state.filtersPicture}
              onChange={this.handleChangePicture}
            />
            <Modal.Footer>
              <Button onClick={this.handleClose}>Fermer</Button>
              <Button bsStyle="primary" onClick={this.addCar}>
                Ajouter
              </Button>
            </Modal.Footer>
          </Modal.Dialog>
        </div>
        <div
          className="static-modal"
          style={{ display: this.state.showResa ? 'inline' : 'none' }}
        >
          <Modal.Dialog>
            <Modal.Header>
              <Modal.Title>Ajouter une réservation</Modal.Title>
            </Modal.Header>
            <input type="text" placeholder="date1" />
            <input type="text" placeholder="date2" />
            <input type="text" placeholder="Status" />
            <input type="text" placeholder="Voiture" />
            <Modal.Footer>
              <Button onClick={this.handleCloseResa}>Fermer</Button>
              <Button bsStyle="primary" onClick={this.addCar}>
                Ajouter
              </Button>
            </Modal.Footer>
          </Modal.Dialog>
        </div>
        <div
          id="filters"
          style={{ opacity: !this.state.show ? '1' : '0.05' }}
        ></div>
        <div id="carslist" style={{ opacity: !this.state.show ? '1' : '0.05' }}>
          <Button bsStyle="primary" onClick={this.handleShow} id="addButton">
            Ajouter un statut
          </Button>
          <div>
            {this.state.cars.length} statuts trouvés{' '}
            <Button id="searchButton" onClick={this.search}>
              Rechercher
            </Button>
            <FormGroup controlId="formControlsSelect" id="sortBar">
              <FormControl componentClass="select" placeholder="select">
                <option value="modele">Modèle</option>
                <option value="agence">Agence</option>
                <option value="annee">Année</option>
                <option value="poids">Poids</option>
                <option value="puissance">Puissance</option>
                <option value="longueur">Longueur</option>
                <option value="largeur">Largeur</option>
                <option value="hauteur">Hauteur</option>
              </FormControl>
            </FormGroup>
          </div>

          {this.tabCars()}
        </div>
      </div>
    );
  }

  handleChangeModel(event) {
    this.setState({ filtersModel: event.target.value });
  }

  handleChangeYearA(event) {
    this.setState({ filtersYearA: event.target.value });
  }

  handleChangeYearB(event) {
    this.setState({ filtersYearB: event.target.value });
  }

  handleChangePowerMin(event) {
    this.setState({ filtersPowerMin: event.target.value });
  }

  handleChangePowerMax(event) {
    this.setState({ filtersPowerMax: event.target.value });
  }

  handleChangeWidth(event) {
    this.setState({ filtersWidth: event.target.value });
  }

  handleChangeLength(event) {
    this.setState({ filtersLength: event.target.value });
  }

  handleChangeHeight(event) {
    this.setState({ filtersHeight: event.target.value });
  }

  handleChangePoids(event) {
    this.setState({ filtersPoids: event.target.value });
  }

  handleChangeImmat(event) {
    this.setState({ filtersImmat: event.target.value });
  }

  handleChangePicture(event) {
    this.setState({ filtersPicture: event.target.value });
  }
}
